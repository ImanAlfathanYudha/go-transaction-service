package error

import "github.com/sirupsen/logrus"

func WrapError(err error) error {
	logrus.Errorf("Error %v:", err)
	return err
}
