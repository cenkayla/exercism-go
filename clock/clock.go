package clock

import "fmt"

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	// Saati dakikaya çevirip dakika ile toplayıp, toplam dakikayı buluyoruz ve 1 gün(1440 dakika) ile modunu alıyoruz
	t := hour*60 + minute
	t %= 24*60
	// Toplam dakika negatif ise geriye doğru sayıyoruz.
	if t < 0 {
		t += 24*60
	}
	// Toplam dakikayı saate ve dakikaya çeviriyoruz
	return Clock{t / 60 % 24, t % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(c.h, c.m-minutes)
}
