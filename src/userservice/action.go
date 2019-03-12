package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gstaad/src/userservice/cognito"
	"gstaad/src/userservice/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	userpb "gstaad/src/userservice/pb"

	"github.com/aws/aws-sdk-go/aws/awserr"
	cogp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	empty "github.com/golang/protobuf/ptypes/empty"
)

const (
	jwtpName = "x-jwt-payload"
)

func (s *server) Signup(c context.Context, in *userpb.SignupRequest) (*userpb.MutationReply, error) {
	s.log.Infof("[Signup] email=%s", in.Email)

	er := s.cc.authSvc.Signup(in.Email, in.Password)
	if er != nil {
		if er, ok := er.(awserr.Error); ok {
			switch er.Code() {
			case cogp.ErrCodeUserNotFoundException:
				return nil, status.Errorf(codes.NotFound, "Already registered email: %s", in.Email)
			case cogp.ErrCodeInvalidParameterException:
				s.log.Warn(er)
				return nil, status.Error(codes.InvalidArgument, "Invalid email or password")
			}
		}

		s.log.Error(er)
		return nil, status.Error(codes.Internal, "Sorry, there's something wrong.")
	}
	return &userpb.MutationReply{Result: true}, nil
}

func (s *server) Login(c context.Context, in *userpb.LoginRequest) (*userpb.LoginReply, error) {
	s.log.Infof("[Login] email=%s", in.Email)

	tk := &cognito.TokenPayload{}
	er := s.cc.authSvc.Login(tk, in.Email, in.Password)
	if er != nil {
		if er, ok := er.(awserr.Error); ok {
			switch er.Code() {
			case cogp.ErrCodeInvalidParameterException, cogp.ErrCodeUserNotFoundException:
				s.log.Warn(er)
				return nil, status.Error(codes.InvalidArgument, "Invalid email or password")
			}
		}

		s.log.Error(er)
		return nil, status.Error(codes.Internal, "Sorry, there's something wrong.")
	}

	return &userpb.LoginReply{
		Token: &userpb.Token{
			AccessToken: tk.AccessToken,
			ExpiresIn:   tk.ExpiresIn,
		},
	}, nil
}

func (s *server) Profile(c context.Context, in *empty.Empty) (*userpb.ProfileReply, error) {
	p := new(types.JWTPayload)
	if er := s.parseJWTPayload(p, c); er != nil {
		s.log.Warnf("Invalid jwt payload: %s", er.Error())
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	s.log.Infof("[Profile]: %s", p.Email)

	cnt := s.postCount(withJWTP(c))
	return &userpb.ProfileReply{
		User: &userpb.User{
			Email: p.Email,
		},
		PostCount: cnt,
	}, nil
}

func (s *server) Unregister(c context.Context, in *userpb.UnregisterRequest) (*userpb.MutationReply, error) {
	p := new(types.JWTPayload)
	if er := s.parseJWTPayload(p, c); er != nil {
		s.log.Warnf("Invalid jwt payload: %s", er.Error())
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	s.log.Infof("[Unregister]: %s", p.Email)

	er := s.cc.authSvc.Unregister(in.Token)
	if er != nil {
		s.log.Error(er)
		return nil, status.Error(codes.Internal, "Sorry, there's something wrong.")
	}
	return &userpb.MutationReply{Result: true}, nil
}

func (s *server) postCount(c context.Context) (cnt int32) {
	c, cancel := context.WithTimeout(c, time.Millisecond*100)
	defer cancel()

	// just a sample to show how to use compressor
	// actually no need to compress for this
	r, er := s.cc.postSvc.Count(c, &empty.Empty{}, grpc.UseCompressor(gzip.Name))
	if er != nil {
		s.log.Errorf("Fail to connect postservice: %v", er)
	} else {
		cnt = r.Count
	}
	return
}

func (s *server) parseJWTPayload(p *types.JWTPayload, ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if b, ok := md[jwtpName]; ok && len(b) > 0 {
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

func withJWTP(c context.Context) context.Context {
	if md, ok := metadata.FromIncomingContext(c); ok {
		if bb, ok := md[jwtpName]; ok {
			for _, b := range bb {
				c = metadata.AppendToOutgoingContext(c, jwtpName, b)
			}
		}
	}
	return c
}
