proto:
	protoc --proto_path=MyGraderList/auth --go_out=. --go-grpc_out=require_unimplemented_servers=false:. auth.proto
	protoc --proto_path==MyGraderList/backend/course --go_out=. --go-grpc_out=require_unimplemented_servers=false:. course.proto
	protoc --proto_path=MyGraderList/backend/emoji --go_out=. --go-grpc_out=require_unimplemented_servers=false:. emoji.proto
	protoc --proto_path=MyGraderList/backend/like --go_out=. --go-grpc_out=require_unimplemented_servers=false:. like.proto
	protoc --proto_path=MyGraderList/backend/problem --go_out=. --go-grpc_out=require_unimplemented_servers=false:. problem.proto
	protoc --proto_path=MyGraderList/backend/rating --go_out=. --go-grpc_out=require_unimplemented_servers=false:. rating.proto
	protoc --proto_path=MyGraderList/backend/user --go_out=. --go-grpc_out=require_unimplemented_servers=false:. user.proto
