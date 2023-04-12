type AddService struct {
	logx.Logger
	dao *dao.Dao
}

func NewAddService(logger logx.Logger, dao *dao.Dao) *AddService {
	return &AddService{
		Logger: logger,
		dao:    dao,
	}
}

func (svc *AddService) Add(ctx context.Context, req types.AddRequest) (*types.AddReply, error) {
	err := svc.dao.Insert(&model.User{
		Name:   req.Name,
		Email:  req.Email,
		DOB:    req.DOB,
		Mobile: req.Mobile,
	})
	if err != nil {
		svc.Error(err)
		return nil, err
	}

	return &types.AddReply{
		Code:    200,
		Message: fmt.Sprintf("record saved successfully for %s", req.Name),
	}, nil
}
