# influxDB-hello-world
 
### Write and query time series data using InfluxDB

<ol>
 <li>Create a bucket and access token in the influxDB UI found under the Getting Started link of the <a href="https://www.influxdata.com/products/influxdb-cloud/serverless">InfluxDB Cloud Serverless</a> webpage</li>
 <li>Clone influxDB-hello-world repository locally</li>
 <li>Set variables
  <ol>
   <li>Set enviroment variables (token and database name) locally (Unix):
    
    export INFLUXDB_TOKEN=<influx-token>
    export INFLUX_DATABASE=<influx-bucket-name>
   </li>
   <li>Change the timestamps in the write.go lines line protocol records to be within the time range configured for your bucket</li>
   <li>Change the host URL in the client initalization found in both the write.go and query.go files</li>
   <li>Run WriteLineProtocol() before trying to query the data in main()</li>
  </ol>
 </li>
 <li>Install the packages listed in imports, build the influxdb_go_client module, and execute the main() function (Unix):
  
    go mod tidy && go build && go run influxdb_go_client
 </li>
</ol>
