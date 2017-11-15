package lib

import (
)

type Cache interface {}

type Redis struct {
	rc interface{}
}