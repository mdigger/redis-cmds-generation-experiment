package cmd

import (
	"testing"
)

func BenchmarkCompleted(b *testing.B) {
	b.Run("via interface", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := Completed(completed{
				command: newCommand("GET", "TEST"),
				ctags:   ctagReadOnly,
			})

			s := c.String()
			_ = s
		}
	})

	b.Run("via interface point", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := Completed(&completed{
				command: newCommand("GET", "TEST"),
				ctags:   ctagReadOnly,
			})
			s := c.String()
			_ = s
		}
	})

	b.Run("direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			z := completed{
				command: newCommand("GET", "TEST"),
				ctags:   ctagReadOnly,
			}

			s := z.String()
			_ = s
		}
	})

	b.Run("direct point", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			z := &completed{
				command: newCommand("GET", "TEST"),
				ctags:   ctagReadOnly,
			}

			s := z.String()
			_ = s
		}
	})
}
