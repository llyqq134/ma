for i in {1..10}; do 
    echo "Request $i:"
    curl -s http://localhost:30080 | grep "Welcome to nginx" | head -1
done
