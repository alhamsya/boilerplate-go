#create proxy
toxiproxy-cli create -l 127.0.0.1:20101 -u 127.0.0.1:6379 redis

#add toxic
toxiproxy-cli toxic add -t latency -a latency=1000 redis

#remove toxic
toxiproxy-cli toxic remove -n latency_downstream redis

#update toxic
toxiproxy-cli toxic update -n latency_downstream -a latency=3000 redis

#toogle toxic
toxiproxy-cli toggle redis