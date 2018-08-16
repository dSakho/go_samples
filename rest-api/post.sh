#!/bin/bash
# Script that creates people
echo "Creating a new person"
echo ""
curl -H "Content-Type: application/json" -X POST -d '{"Firstname": "Daouda", "Lastname" : "Sakho", "Address": {"City": "Jackson Heights", "State": "New York"}}' http://localhost:8000/people

echo "Creating Ronald McDonald"
echo ""
curl -H "Content-Type: application/json" -X POST -d '{"Firstname": "Ronald", "Lastname" : "McDonald", "Address": {"City": "The", "State": "World"}}' http://localhost:8000/people

echo "Creating Iron Man"
echo ""
curl -H "Content-Type: application/json" -X POST -d '{"Firstname": "Tony", "Lastname" : "Stark", "Address": {"City": "New York", "State": "NY"}}' http://localhost:8000/people

echo "Creating Captain America"
echo ""
curl -H "Content-Type: application/json" -X POST -d '{"Firstname": "Steve", "Lastname" : "Rogers", "Address": {"City": "Brooklyn", "State": "NY"}}' http://localhost:8000/people

                                                                                                                                                                                    