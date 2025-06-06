# TidyBeaver
## WIP - Log Aggregator

Development Steps
- [X] 1. Watch a directory and parse logs in real-time
- [X] 2. Save structured logs (JSON) into local files
- [X] 3. Send logs via HTTP POST and store them
- [X] 4. Query logs by service or severity
- [X] 5. Add structured logging from sample Go microservices
- [X] 6. Accept logs in multiple formats (plain text, JSON)
- [X] 8. Create pluggable parsers for different log sources
- [ ] 9. A CloudWatch logs ingester
- [ ] 10. Push alerts to Slack
- [X] 11. Build a web dashboard using Go (with gin + html/template) that queries your aggregator API.
- [ ] 12. Unit testing with testing package
- [X] 13. Logging best practices (structured logs, correlation IDs)
- [X] 16. Add to README: architecture, instructions and screenshots
- [X] 17. Deploy on docker container

Business Rules
- [X] 1. Get and or Receive logs from multiple sources
- [X] 2. Save logs in a AWS S3 bucket
- [X] 3. Index/Parse the saved logs
- [X] 4. Be able to analyze logs through a UI