package main

import (
  "context"
  "os"
  "fmt"
  "log"

  "github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

// Write line protocol data to InfluxDB
func WriteLineProtocol() error {
  // INFLUX_TOKEN is an environment variable you assigned to your
  // API WRITE token value.
  token := os.Getenv("INFLUX_TOKEN")
  database := os.Getenv("INFLUX_DATABASE")

  // Initialize a client with URL and token
  client, err := influxdb3.New(influxdb3.ClientConfig{
    Host:     "https://us-east-1-1.aws.cloud2.influxdata.com/",
    Token:    token,
    Database: database,
  })

  // Close the client when the function returns.
  defer func(client *influxdb3.Client) {
    err := client.Close()
    if err != nil {
      panic(err)
    }
  }(client)

  // Define line protocol records to write.
  // Use a raw string literal (denoted by backticks)
  // to preserve backslashes and prevent interpretation
  // of escape sequences--for example, escaped spaces in tag values.
  lines := [...]string{
    `home,room=Living\ Room temp=21.1,hum=35.9,co=0i 1704528000000000000`,
    `home,room=Kitchen temp=21.0,hum=35.9,co=0i 1704528000000000000`,
    `home,room=Living\ Room temp=21.4,hum=35.9,co=0i 1704531600000000000`,
    `home,room=Kitchen temp=23.0,hum=36.2,co=0i 1704531600000000000`,
    `home,room=Living\ Room temp=21.8,hum=36.0,co=0i 1704535200000000000`,
    `home,room=Kitchen temp=22.7,hum=36.1,co=0i 1704535200000000000`,
    `home,room=Living\ Room temp=22.2,hum=36.0,co=0i 1704538800000000000`,
    `home,room=Kitchen temp=22.4,hum=36.0,co=0i 1704538800000000000`,
    `home,room=Living\ Room temp=22.2,hum=35.9,co=0i 1704542400000000000`,
    `home,room=Kitchen temp=22.5,hum=36.0,co=0i 1704542400000000000`,
    `home,room=Living\ Room temp=22.4,hum=36.0,co=0i 1704546000000000000`,
    `home,room=Kitchen temp=22.8,hum=36.5,co=1i 1704546000000000000`,
    `home,room=Living\ Room temp=22.3,hum=36.1,co=0i 1704549600000000000`,
    `home,room=Kitchen temp=22.8,hum=36.3,co=1i 1704549600000000000`,
    `home,room=Living\ Room temp=22.3,hum=36.1,co=1i 1704553200000000000`,
    `home,room=Kitchen temp=22.7,hum=36.2,co=3i 1704553200000000000`,
    `home,room=Living\ Room temp=22.4,hum=36.0,co=4i 1704556800000000000`,
    `home,room=Kitchen temp=22.4,hum=36.0,co=7i 1704556800000000000`,
    `home,room=Living\ Room temp=22.6,hum=35.9,co=5i 1704560400000000000`,
    `home,room=Kitchen temp=22.7,hum=36.0,co=9i 1704560400000000000`,
    `home,room=Living\ Room temp=22.8,hum=36.2,co=9i 1704564000000000000`,
    `home,room=Kitchen temp=23.3,hum=36.9,co=18i 1704564000000000000`,
    `home,room=Living\ Room temp=22.5,hum=36.3,co=14i 1704567600000000000`,
    `home,room=Kitchen temp=23.1,hum=36.6,co=22i 1704567600000000000`,
    `home,room=Living\ Room temp=22.2,hum=36.4,co=17i 1704571200000000000`,
    `home,room=Kitchen temp=22.7,hum=36.5,co=26i 1704571200000000000`,
  }

  // Iterate over the lines array and write each line
  // separately to InfluxDB
  for _, record := range lines {
    err = client.Write(context.Background(), []byte(record))
    if err != nil {
      log.Fatalf("Error writing line protocol: %v", err)
    }
  }

  if err != nil {
    panic(err)
  }

  fmt.Println("Data has been written successfully.")
  return nil
}