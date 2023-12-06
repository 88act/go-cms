// Code generated by goctl. DO NOT EDIT!
// Source: basic.proto

package basic

import (
	"context"

	"go-cms/app/basic/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFileReq     = pb.AddFileReq
	AddFileResp    = pb.AddFileResp
	BasicFile      = pb.BasicFile
	CaptchaReq     = pb.CaptchaReq
	CaptchaResp    = pb.CaptchaResp
	FileDetailReq  = pb.FileDetailReq
	FileDetailResp = pb.FileDetailResp
	FileInfo       = pb.FileInfo
	FileListReq    = pb.FileListReq
	FileListResp   = pb.FileListResp
	SendCodeReq    = pb.SendCodeReq
	SendCodeResp   = pb.SendCodeResp
	VerifyCodeReq  = pb.VerifyCodeReq
	VerifyCodeResp = pb.VerifyCodeResp

	Basic interface {
		// 图形码
		Captcha(ctx context.Context, in *CaptchaReq, opts ...grpc.CallOption) (*CaptchaResp, error)
		// 验证码
		SendCode(ctx context.Context, in *SendCodeReq, opts ...grpc.CallOption) (*SendCodeResp, error)
		// 图形码
		VerifyCaptcha(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeResp, error)
		// 图形码
		VerifyCode(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeResp, error)
		// FileDetail //文件详情
		FileDetail(ctx context.Context, in *FileDetailReq, opts ...grpc.CallOption) (*FileDetailResp, error)
		// FileList 文件列表
		FileList(ctx context.Context, in *FileListReq, opts ...grpc.CallOption) (*FileListResp, error)
		// addFile 新增文件
		AddFile(ctx context.Context, in *AddFileReq, opts ...grpc.CallOption) (*AddFileResp, error)
	}

	defaultBasic struct {
		cli zrpc.Client
	}
)

func NewBasic(cli zrpc.Client) Basic {
	return &defaultBasic{
		cli: cli,
	}
}

// 图形码
func (m *defaultBasic) Captcha(ctx context.Context, in *CaptchaReq, opts ...grpc.CallOption) (*CaptchaResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.Captcha(ctx, in, opts...)
}

// 验证码
func (m *defaultBasic) SendCode(ctx context.Context, in *SendCodeReq, opts ...grpc.CallOption) (*SendCodeResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.SendCode(ctx, in, opts...)
}

// 图形码
func (m *defaultBasic) VerifyCaptcha(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.VerifyCaptcha(ctx, in, opts...)
}

// 图形码
func (m *defaultBasic) VerifyCode(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.VerifyCode(ctx, in, opts...)
}

// FileDetail //文件详情
func (m *defaultBasic) FileDetail(ctx context.Context, in *FileDetailReq, opts ...grpc.CallOption) (*FileDetailResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.FileDetail(ctx, in, opts...)
}

// FileList 文件列表
func (m *defaultBasic) FileList(ctx context.Context, in *FileListReq, opts ...grpc.CallOption) (*FileListResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.FileList(ctx, in, opts...)
}

// addFile 新增文件
func (m *defaultBasic) AddFile(ctx context.Context, in *AddFileReq, opts ...grpc.CallOption) (*AddFileResp, error) {
	client := pb.NewBasicClient(m.cli.Conn())
	return client.AddFile(ctx, in, opts...)
}
