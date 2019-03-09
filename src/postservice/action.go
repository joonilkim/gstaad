package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	postpb "gstaad/src/postservice/pb"
	"gstaad/src/postservice/types"
)

func (s *server) Create(c context.Context, in *postpb.CreateRequest) (*postpb.MutationReply, error) {
	p := new(types.JWTPayload)
	if er := s.parseJWTPayload(p, c); er != nil {
		s.log.Warnf("Invalid jwt payload: %s", er.Error())
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	s.log.Infof("[CreatePost]: email=%s", p.Email)

	return &postpb.MutationReply{Result: true}, nil
}

func (s *server) All(c context.Context, in *empty.Empty) (*postpb.PostsReply, error) {
	p := new(types.JWTPayload)
	if er := s.parseJWTPayload(p, c); er != nil {
		s.log.Warnf("Invalid jwt payload: %s", er.Error())
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	s.log.Infof("[AllPost]: email=%s", p.Email)

	author := postpb.User{
		Email: "aa@aa.com",
	}

	posts := []*postpb.Post{
		&postpb.Post{
			Id:        "post1",
			Author:    &author,
			Content:   "post1",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
		&postpb.Post{
			Id:        "post2",
			Author:    &author,
			Content:   "post2",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
	}
	return &postpb.PostsReply{
		Items: posts,
	}, nil
}

func (s *server) Count(c context.Context, in *empty.Empty) (*postpb.CountReply, error) {
	p := new(types.JWTPayload)
	if er := s.parseJWTPayload(p, c); er != nil {
		s.log.Warnf("Invalid jwt payload: %s", er.Error())
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	s.log.Infof("[PostCount]: email=%s", p.Email)

	return &postpb.CountReply{
		Count: 10,
	}, nil
}

func (s *server) parseJWTPayload(p *types.JWTPayload, ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if b, ok := md["x-jwt-payload"]; ok && len(b) > 0 {
			return s.decodeJWTPayload(p, b[0])
		}
		return fmt.Errorf("Invalid jwt payload: %v", md)
	}
	return fmt.Errorf("No metadata Found")
}

func (s *server) decodeJWTPayload(p *types.JWTPayload, b string) error {
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b))
	return json.NewDecoder(r).Decode(p)
}
