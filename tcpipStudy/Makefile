build:
	docker-compose build --no-cache --progress plain
up:
	docker-compose up -d
down:
	docker-compose down
remove:
	docker-compose down --remove-orphans
tuntap:
	ip tuntap add mode tun dev tun0 &&\
	ip link set tun0 up &&\
	ip addr add 10.0.0.1/24 dev tun0

tundetouch:
	ip addr del 10.0.0.1/24 dev tun0 &&\
	ip link delete tgo un0 &&\
	ip tuntap del mode tun dev tun0

capture:
	tcpdump -i tun0

curl:
	curl --interface tun0 http://10.0.0.2/
