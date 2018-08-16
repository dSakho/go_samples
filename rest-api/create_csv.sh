#!/bin/bash
# Creating CSV of people
echo "Creating CSV file of people"
echo ""
curl -H "Content-Type: application/json" -X POST http://localhost:8000/people