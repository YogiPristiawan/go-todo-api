package validators

import (
	"fmt"
	"go_todo_api/src/shared/validators"
	"go_todo_api/src/todo/dto"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TodoValidatorTestSuite struct {
	suite.Suite
	todoValidator TodoValidator
}

func (s *TodoValidatorTestSuite) SetupSuite() {
	s.todoValidator = NewTodoValidator(validators.NewValidator())
}

func TestTodoValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(TodoValidatorTestSuite))
}

func (s *TodoValidatorTestSuite) TestNewTodoValidator() {
	s.Run("It should properly instantiate todoValidator", func() {
		// arrange
		validator := validators.NewValidator()
		var expect *todoValidator

		// action
		todoV := NewTodoValidator(validator)

		// assert
		s.Assert().IsTypef(expect, todoV, "Should return type of %T", expect)
		s.Assert().Implementsf((*TodoValidator)(nil), todoV, "Should implement %T", (*TodoValidator)(nil))
	})
}

func (s *TodoValidatorTestSuite) TestValidateStore() {
	type test struct {
		title     string
		param     dto.StoreTodoRequest
		expectErr *validators.ValidatorError
	}

	s.Run("It should validate todo", func() {
		// NEGATIVE
		// arrange
		negatives := []test{
			{
				title: "It should return an error if todo is not provided",
				param: dto.StoreTodoRequest{
					Date:       "2022-01-01",
					IsFinished: false,
				},
			},
			{
				title: "It should return an error if todo has a length greater than 255",
				param: dto.StoreTodoRequest{
					Todo:       "*%t#8p!UDmcymTzd?3nc8K:-TU.%n=#/9*iWU2a#6Y!%wDv&G{qAUZm#Luvq]dCkSgSQ:mS2j.QXYwruvPtwwnJp?*W@[F!xPBC##nDmJ2}9%h6,KNq3Hv%)PU7fS%RX!3/z$vd.t@r/%ZkTm6!z*i5DZVy%abR=Qbvzyu.AT/V(n7[Y@H.#c}@?yvbn[yLf7hV6QF&ZiX5H4?ih,&Yh+.HxNk}7.4*E{=8v98NJ=g)R{Q5nuE5E%B57:](}3e/KK3M{(iGzH5ynKm3fbVpfBQz8dA}a#b_pKd]b(jg7-]_D=rLdua%M[?TY@J,F(qUvpWeqt6+8m!jt9p[Brj8LH}L%r$72Wit!NRwm;NN(@k[n3iu=.k8&S%?r#}]NA2Ccce$KL,kWe.k%9v$BW)5$F3#ckVr8TFRr9K4+h[A?[J6[rt}:T82ngWWfHWu%=[Aiy#)FH=5EEy!d)un4/qXZx,N@d.K2AX_*3u[[SJ@fnB[[Cu!{rjG2",
					Date:       "2022-01-01",
					IsFinished: false,
				},
			},
		}

		// action & assert
		for _, test := range negatives {
			err := s.todoValidator.ValidateStore(test.param)

			message := "%s"
			message += "\nTodo: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message, test.param.Todo,
			)
		}

		// POSITIVE
		// arrange
		positives := []test{
			{
				title: "Should not return an error if the given todo is less than 255 length",
				param: dto.StoreTodoRequest{
					Todo:       "hari ini aku mau ngapain",
					Date:       "2022-01-01",
					IsFinished: false,
				},
			},
		}

		// action & assert
		for _, test := range positives {
			err := s.todoValidator.ValidateStore(test.param)
			fmt.Println(err)

			message := "%s"
			message += "\nTodo: %s"

			s.Assert().Nilf(
				err,
				message, test.param.Todo,
			)
		}
	})

	s.Run("It should validate date", func() {
		// arrange
		negatives := []test{
			{
				title: "Should return an error if the date is not provided",
				param: dto.StoreTodoRequest{
					Todo:       "hari ini aku mau ngapain",
					IsFinished: true,
				},
			},
			{
				title: "Should return an error if the given date format is invalid",
				param: dto.StoreTodoRequest{
					Todo:       "hari ini aku mau ngapain",
					Date:       "2022",
					IsFinished: true,
				},
			},
			{
				title: "Should return an error if the given date format is invalid",
				param: dto.StoreTodoRequest{
					Todo:       "hari ini aku mau ngapain",
					Date:       "09",
					IsFinished: true,
				},
			},
			{
				title: "Should return an error if the given date format is invalid",
				param: dto.StoreTodoRequest{
					Todo:       "hari ini aku mau ngapain",
					Date:       "January",
					IsFinished: true,
				},
			},
		}

		// action & assert
		for _, test := range negatives {
			err := s.todoValidator.ValidateStore(test.param)

			message := "%s"
			message += "\nDate: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message, test.param.Date,
			)
		}

		// POSITIVE
		// arrange
		positives := []test{
			{
				title: "Should not return an error if the given date format is valid",
				param: dto.StoreTodoRequest{
					Todo:       "hari ini aku mau ngapain",
					Date:       "2022-01-01",
					IsFinished: true,
				},
			},
		}

		// action & assert
		for _, test := range positives {
			err := s.todoValidator.ValidateStore(test.param)

			message := "%s"
			message += "\nDate: %s"

			s.Assert().Nilf(
				err,
				message, test.title, test.param.Date,
			)
		}
	})
}

func (s *TodoValidatorTestSuite) TestValidateDetail() {
	// arange
	type test struct {
		title     string
		param     dto.DetailTodoRequest
		expectErr *validators.ValidatorError
	}

	s.Run("It should validate id", func() {
		// arrange
		negatives := []test{
			{
				title: "Should return an error if id is not provided",
				param: dto.DetailTodoRequest{},
			},
		}

		// action & assert
		for _, test := range negatives {
			err := s.todoValidator.ValidateDetail(test.param)

			message := "%s"
			message += "\nId: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message, test.title, test.param.Id,
			)
		}

	})
}
