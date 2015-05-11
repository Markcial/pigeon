package Pigeon_test

import (
	. "github.com/Markcial/pigeon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agenda", func() {

    var agenda Agenda

    BeforeEach(func() {
       agenda = NewAgenda()
    })

    Describe("Agenda contact manipulation", func() {

        Context("Empty agenda", func() {
            It("should have no contacts by default", func() {
                Expect(len(agenda.All())).To(Equal(0))
            })
        })

        Context("Contact manipulation", func() {
            It("should store contacts successfully", func() {
                agenda.Add(NewContact())
                Expect(len(agenda.All())).To(Equal(1))
                agenda.Add(NewContact())
                Expect(len(agenda.All())).To(Equal(2))
            })
        })
    })
})
