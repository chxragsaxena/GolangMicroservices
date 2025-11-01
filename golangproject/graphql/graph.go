package main 

type Server struct{
	accountClient
	catalogClient
	orderClient

}


func NewGraphQLServer(accountUrl, catalogUrl, orderUrl, string) (*Server, error){
	accountClient , err := account.NewClient(accountUrl)
	if err != nil{
		return nil, err
	}

	catalogClient, err := catalog.NewClient(catalogUrl)
	if err != nil{
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)
	if err != nil{
		accountClient.Close()
		catalogClient.Close()
		return nul, err
	}

	return &Server{
		accountClient,
		catalogClient,
		orderClient,
	}, nil
}

func (s *Server) Mutation() MutationResovler{
	return &mutationResolver{
		server: s,
	}
}


func (s *Server) Query() QueryResolver{
	return &queryResolver{
		server: s,
	}
}

func (s *Server) Account() AccountResolver{
	return &accountResolver{
		server: s,
	}
}

func (s *Server) ToExecuatableSchema() graphQL.ExecuatableSchema{
    return NewExecutableSchema(Config{

	Resolvers: s,	
})
}