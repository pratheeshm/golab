# Go Map Performance Testing

This project demonstrates the performance of map operations in Go by measuring lookup, insertion, and deletion times. It includes Dockerfiles for two different versions of Go (1.23 and 1.24) to facilitate testing across versions.

## Project Structure

```
go-map
├── Dockerfile-go1.23
├── Dockerfile-go1.24
├── main.go
├── go.mod
└── README.md
```

## Dockerfiles

### Dockerfile-go1.23

This Dockerfile sets up a Docker image using Go version 1.23. It installs the necessary dependencies, copies the source code, and specifies how to run the application.

### Dockerfile-go1.24

This Dockerfile is similar to the one for Go 1.23 but uses Go version 1.24. It ensures that the application runs in the latest Go environment.

## Building the Docker Images

To build the Docker images, run the following commands in the project directory:

```bash
docker build -t go-map:1.23 -f Dockerfile-go1.23 .
docker build -t go-map:1.24 -f Dockerfile-go1.24 .
```

## Running the Application

After building the images, you can run the application using the following commands:

```bash
docker run --rm go-map:1.23
docker run --rm go-map:1.24
```

### Performance Comparison

Lookup/Insertion/Deletion time is the total time taken to complete the operation on N keys

#### map[int]int 1M

| Operation | Go 1.23                      | Go 1.24                      |
| --------- | ---------------------------- | ---------------------------- |
| Lookup    | 411.67ms, 422.07ms, 421.51ms | 291.40ms, 277.62ms, 324.95ms |
| Insertion | 239.31ms, 239.54ms, 249.33ms | 149.78ms, 172.36ms, 152.76ms |
| Deletion  | 68.86ms, 71.40ms, 66.96ms    | 99.90ms, 99.51ms, 106.57ms   |

#### map[string]string 1M

| Operation | Go 1.23                      | Go 1.24                      |
| --------- | ---------------------------- | ---------------------------- |
| Lookup    | 1.63s, 1.62s, 1.63s          | 2.01s, 1.97s, 1.98s          |
| Insertion | 582.42ms, 602.09ms, 599.01ms | 515.58ms, 501.34ms, 521.62ms |
| Deletion  | 318.17ms, 303.78ms, 311.75ms | 245.16ms, 248.00ms, 261.74ms |

#### map[string]string 10M

| Operation | Go 1.23             | Go 1.24             |
| --------- | ------------------- | ------------------- |
| Lookup    | 2.67s, 2.49s, 2.48s | 3.35s, 3.29s, 3.43s |
| Insertion | 6.89s, 6.29s, 6.29s | 6.46s, 6.47s, 6.52s |
| Deletion  | 3.49s, 3.65s, 3.51s | 3.08s, 3.16s, 3.28s |

#### map[string]string 100M

| Operation | Go 1.23  | Go 1.24  |
| --------- | -------- | -------- |
| Lookup    | 28.56s   | 42.43s   |
| Insertion | 1m10.47s | 1m18.50s |
| Deletion  | 39.81s   | 46.46s   |

### Final Summary

From the performance comparison, we can observe that:
- Go 1.24 generally shows improved performance for `map[int]int` operations, especially for lookup and insertion.
- For `map[string]string`, Go 1.23 performs better in lookup operations, while Go 1.24 shows improvements in insertion and deletion times for smaller datasets.
- As the dataset size increases to 100M, Go 1.23 outperforms Go 1.24 in all operations.

