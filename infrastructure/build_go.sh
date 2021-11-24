#!/bin/bash
cd backend/
go build -o ../srv
cd ..
sudo ./srv
