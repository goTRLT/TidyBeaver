# TidyBeaver
## WIP - Log Aggregator

Development Steps
- [X] 1. Watch a directory and parse logs in real-time
- [X] 2. Save structured logs (JSON) into local files
- [X] 3. Send logs via HTTP POST and store them
- [ ] 4. Query logs by service or severity
- [ ] 5. Add structured logging from sample Go microservices
- [X] 6. Accept logs in multiple formats (plain text, JSON)
- [ ] 7. Tail log files (tail lib or os/exec with tail -f)
- [X] 8. Create pluggable parsers for different log sources
- [ ] 9. A CloudWatch logs ingester
- [ ] 10. Push alerts to Slack
- [ ] 11. Build a web dashboard using Go (with gin + html/template) that queries your aggregator API.
- [ ] 12. Unit testing with testing package
- [X] 13. Logging best practices (structured logs, correlation IDs)
- [ ] 14. Use gRPC for service-to-service log transfer
- [ ] 15. Use Kafka for log transfer
- [ ] 16. Add to README: architecture, instructions and screenshots
- [ ] 17. Deploy on docker container

Business Rules
- [X] 1. Get and or Receive logs from multiple sources
- [ ] 2. Save logs in a AWS S3 bucket
- [ ] 3. Index/Parse the saved logs
- [ ] 4. Be able to analyze logs through a UI