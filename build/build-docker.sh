docker build  -t hobord/infra2:session -f build/session/Dockerfile .
docker build  -t hobord/infra2:redirect -f build/redirect/Dockerfile .
docker build  -t hobord/infra2:infra -f build/infra/Dockerfile .