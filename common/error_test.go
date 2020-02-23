package common_test

import "testing"

func TestError(t *testing.T) {

}

// var _ = Describe("RwG Error Tests", func() {

// 	Context("When creating an error", func() {
// 		It("should create", func() {
// 			err := NewAPIError("")
// 			Expect(err.ErrorType).To(Equal(APIError))
// 			Expect(err.ErrorType.Name()).To(Equal("APIError"))
// 			Expect(err.Error()).To(Equal("API Error: "))

// 			err = NewConfigurationError("")
// 			Expect(err.ErrorType).To(Equal(ConfigError))
// 			Expect(err.ErrorType.Name()).To(Equal("ConfigError"))
// 			Expect(err.Error()).To(Equal("Configuration Error: "))

// 			err = NewDatabaseOperationError("")
// 			Expect(err.ErrorType).To(Equal(DatabaseError))
// 			Expect(err.ErrorType.Name()).To(Equal("DatabaseError"))
// 			Expect(err.Error()).To(Equal("Database Operation Error: "))

// 			err = NewSystemError("")
// 			Expect(err.ErrorType).To(Equal(SystemError))
// 			Expect(err.ErrorType.Name()).To(Equal("SystemError"))
// 			Expect(err.Error()).To(Equal("System Error: "))

// 			err = NewDataError("")
// 			Expect(err.ErrorType).To(Equal(DataError))
// 			Expect(err.ErrorType.Name()).To(Equal("DataError"))
// 			Expect(err.Error()).To(Equal("Data Error: "))

// 			err = NewValidationError("")
// 			Expect(err.ErrorType).To(Equal(ValidationError))
// 			Expect(err.ErrorType.Name()).To(Equal("ValidationError"))
// 			Expect(err.Error()).To(Equal("Validation Error: "))

// 			errs := make([]*Error, 0)
// 			valErr := NewValidationError("missing id")
// 			valErr2 := NewValidationError("missing price")
// 			errs = append(errs, valErr, valErr2)
// 			err = NewValidationConcatError(errs)
// 			Expect(err.ErrorType).To(Equal(ValidationError))
// 			Expect(err.ErrorType.Name()).To(Equal("ValidationError"))
// 			Expect(err.Error()).To(Equal("Validation Error: missing id, missing price,"))

// 			errs = NewSingletonErrorList(valErr)
// 			Expect(len(errs)).To(Equal(1))
// 			Expect(errs[0]).To(Equal(valErr))
// 		})
// 	})
// })
