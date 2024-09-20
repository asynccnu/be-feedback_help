package grpc

import (
	"context"
	"errors"
	pb "github.com/asynccnu/be-api/gen/proto/feedback_help/v1"
	"github.com/asynccnu/be-feedback_help/domain"
	ser "github.com/asynccnu/be-feedback_help/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
)

type FeedbackHelpService struct {
	pb.UnimplementedFeedbackHelpServer
	ser ser.Service
}

func NewFeedbackHelpServiceServer(ser ser.Service) *FeedbackHelpService {
	return &FeedbackHelpService{ser: ser}
}

func (s *FeedbackHelpService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterFeedbackHelpServer(server, s)
}
func (s *FeedbackHelpService) GetQuestions(ctx context.Context, req *pb.EmptyRequest) (*pb.GetQuestionsResponse, error) {
	q, err := s.ser.GetQuestions(ctx)
	if err != nil {
		return nil, err
	}
	var questions []*pb.FrequentlyAskedQuestion
	for _, q := range q {
		questions = append(questions, &pb.FrequentlyAskedQuestion{
			Id:         q.Id,
			Question:   q.Question,
			Answer:     q.Answer,
			ClickTimes: int64(q.ClickTimes),
		})
	}
	return &pb.GetQuestionsResponse{
		Questions: questions,
	}, nil
}
func (s *FeedbackHelpService) FindQuestionByName(ctx context.Context, req *pb.FindQuestionByNameRequest) (*pb.FindQuestionByNameResponse, error) {
	q, err := s.ser.FindQuestionByName(ctx, req.Question)
	if err != nil {
		return nil, err
	}
	var questions []*pb.FrequentlyAskedQuestion
	for _, q := range q {
		questions = append(questions, &pb.FrequentlyAskedQuestion{
			Id:         q.Id,
			Question:   q.Question,
			Answer:     q.Answer,
			ClickTimes: int64(q.ClickTimes),
		})
	}
	//异步记录
	go func() {
		err := s.ser.NoteMoreFeedbackSearch(ctx, domain.EventSearchQuestion{
			Question: req.Question,
		})
		if err != nil {
			log.Println(err)
		}
	}()
	return &pb.FindQuestionByNameResponse{
		Questions: questions,
	}, nil
}
func (s *FeedbackHelpService) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.OperationResponse, error) {
	err := s.ser.CreateQuestion(ctx, domain.FrequentlyAskedQuestion{
		Question: req.Question,
		Answer:   req.Anwser,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OperationResponse{}, nil
}
func (s *FeedbackHelpService) ChangeQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.OperationResponse, error) {
	err := s.ser.ChangeQuestion(ctx, domain.FrequentlyAskedQuestion{
		Id:       req.QuestionId,
		Question: req.Question,
		Answer:   req.Anwser,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OperationResponse{}, nil
}
func (s *FeedbackHelpService) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.OperationResponse, error) {
	err := s.ser.DeleteQuestion(ctx, domain.FrequentlyAskedQuestion{
		Id: req.QuestionId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OperationResponse{}, nil
}
func (s *FeedbackHelpService) NoteQuestion(ctx context.Context, req *pb.NoteQuestionRequest) (*pb.OperationResponse, error) {
	err := s.ser.NoteQuestion(ctx, domain.Question{
		QuestionId: req.QuestionId,
		IfOver:     req.IfOver,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OperationResponse{}, nil
}
func (s *FeedbackHelpService) NoteEventTracking(ctx context.Context, req *pb.NoteEventTrackingRequest) (*pb.OperationResponse, error) {

	if req.Event < 0 || req.Event > 3 {
		return nil, errors.New("无效类型")
	}
	err := s.ser.NoteEventTracking(ctx, domain.EventTracking{
		Event: int8(req.Event),
	})
	if err != nil {
		return nil, err
	}
	return &pb.OperationResponse{}, nil
}
func (s *FeedbackHelpService) NoteMoreFeedbackSearchSkip(ctx context.Context, req *pb.NoteMoreFeedbackSearchSkipRequest) (*pb.OperationResponse, error) {
	err := s.ser.NoteMoreFeedbackSearchSkip(ctx, domain.EventQuestion{
		QuestionId: req.QuestionId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OperationResponse{}, nil
}

type FrequentlyAskedQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Question      string `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Answer        string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
	ClickTimes    int64  `protobuf:"varint,4,opt,name=click_times,json=clickTimes,proto3" json:"click_times,omitempty"`
}
