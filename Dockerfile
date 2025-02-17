# Build stage
FROM golang:1.23-bookworm AS build

WORKDIR /app

# Copy only necessary files
COPY go.mod go.sum ./
RUN go mod download

COPY project ./project
COPY protobuf ./protobuf

# Build the application
WORKDIR /app/project
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/project-binary main.go

# Final stage
FROM gcr.io/distroless/static:nonroot AS final
WORKDIR /app

COPY --from=build /app/project-binary .
CMD [ "./project-binary" ]


#Build stage
#FROM golang:1.23-bookworm AS build

#RUN apt-get update && apt-get install -y busybox && apt-get clean

#WORKDIR /app

# Copy only necessary files
#COPY go.mod go.sum ./
#RUN go mod download

#COPY project ./project
#COPY protobuf ./protobuf

# Build the application
#WORKDIR /app/project
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app/project/project-binary main.go

# Final stage
#FROM gcr.io/distroless/static:nonroot AS final
#WORKDIR /app

#COPY --from=build /app/bin/project /app/project

#CMD [ "./project/project-binary" ]
