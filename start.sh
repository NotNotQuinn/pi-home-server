echo "Starting websub server..."
go run ./websub &
sleep 1
echo "Starting keylogger server (as root)..."
sudo go run ./keylogger &
