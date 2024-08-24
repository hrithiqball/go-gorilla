#!/bin/bash

# Read TESTING.md section on DDoS attack

VEGETA_PATH="$HOME/go/bin/vegeta"

TARGET_URL="http://localhost:8080/user/list"
DURATION="30s"
RATE="1000"

# Run Vegeta attack
echo "GET $TARGET_URL" | $VEGETA_PATH attack -duration=$DURATION -rate=$RATE | $VEGETA_PATH report
