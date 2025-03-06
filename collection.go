package validator

import "github.com/Pashgunt/Validator/internal/contract"

type AssertListValue []contract.ConstraintInterface
type AssertList map[string]AssertListValue

type CollectionAssertsInterface interface {
	Asserts() AssertList
	SetAsserts(asserts AssertList)
}

type CollectionInterface interface {
	RemoveAllAssertsByProperty(property string)
	Exist(property string) bool
	CollectionAssertsInterface
}

type Collection struct {
	asserts AssertList
}

func NewCollection(asserts AssertList) *Collection {
	return &Collection{asserts: asserts}
}

func (c *Collection) Asserts() AssertList {
	return c.asserts
}

func (c *Collection) SetAsserts(asserts AssertList) {
	c.asserts = asserts
}

func (c *Collection) RemoveAllAssertsByProperty(property string) {
	if !c.Exist(property) {
		return
	}

	delete(c.Asserts(), property)
}

func (c *Collection) Exist(property string) bool {
	_, isset := c.Asserts()[property]

	return isset
}
