package main

import (
	pb "InstrumentoM/proto"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var db *sql.DB

type server struct {
	pb.UnimplementedInstrumentServiceServer
}

func (s *server) GetInstrumentinfo(ctx context.Context, req *pb.InstrumentRequest) (*pb.InstrumentResponse, error) {
	var name, ptype, origin string

	query := "select * from instruprd.instrumento where Name=@Name"

	row := db.QueryRowContext(ctx, query, sql.Named("Name", req.Name))
	err := row.Scan(&name, &ptype, &origin)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.InstrumentResponse{
				Name:   "Not Found",
				Type:   "Not Found",
				Origin: "Not Found",
			}, nil
		}
		return nil, err
	}

	return &pb.InstrumentResponse{
		Name:   name,
		Type:   ptype,
		Origin: origin,
	}, nil
}

func (s *server) GetInstrumentList(req *pb.Empty, stream pb.InstrumentService_GetInstrumentListServer) error {
	query := "SELECT * FROM instruprd.instrumento"
	rows, err := db.Query(query)

	if err != nil {
		log.Panic(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name, ptype string
		var origin sql.NullString

		if err := rows.Scan(&name, &ptype, &origin); err != nil {
			log.Panic(err)
			return err
		}

		var originValue string
		if origin.Valid {
			originValue = origin.String
		} else {
			originValue = "Unknown"
		}

		if err := stream.Send(&pb.InstrumentResponse{
			Name:   name,
			Type:   ptype,
			Origin: originValue,
		}); err != nil {
			log.Panic(err)
			return err
		}
	}
	return nil
}

func (s *server) AddInstrumentos(Stream pb.InstrumentService_AddInstrumentosServer) error {
	var count int32
	for {
		req, err := Stream.Recv()
		if err == io.EOF {
			return Stream.SendAndClose(&pb.AddInstrumentResponse{
				Count: count,
			})
		}
		if err != nil {
			log.Panic(err)
			return err
		}

		query := "insert into instruprd.instrumento (Name, Type,Origin)values(@Name,@Type,@Origin)"
		_, err = db.Exec(query,
			sql.Named("Name", req.Name),
			sql.Named("Type", req.Type),
			sql.Named("Origin", req.Origin))

		if err != nil {
			log.Panic(err)
			return err
		}
		count++

	}
}

func (s *server) GetInstrumentosByType(stream pb.InstrumentService_GetInstrumentosByTypeServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			return nil

		}
		if err != nil {
			log.Panic(err)
			return err

		}
		query := "Select * from instruprd.instrumento where Type=@Type"
		rows, err := db.Query(query, sql.Named("Type", req.Type))
		if err != nil {
			log.Panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var name, ptype string
			var origin string

			if err := rows.Scan(&name, &ptype, &origin); err != nil {
				log.Panic(err)
				return err
			}
			if err := stream.Send(&pb.InstrumentResponse{
				Name:   name,
				Type:   ptype,
				Origin: origin,
			}); err != nil {
				log.Panic(err)
				return err

			}

		}
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	s := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, s, port, database)

	db, err = sql.Open("sqlserver", connString) // Aquí se usa la variable global `db`
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInstrumentServiceServer(grpcServer, &server{})

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK")) // Aquí se usa []byte("OK") en lugar de []byte{"OK"}
		})
		log.Println("Starting health check server on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Println("starting server on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
