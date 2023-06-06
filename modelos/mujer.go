package modelos

type Mujer struct {
	Hombre
	EsMadre bool
}

func (m *Mujer) Respirar()      { m.respirando = true }
func (m *Mujer) Comer()         { m.respirando = true }
func (m *Mujer) Pensar()        { m.respirando = true }
func (m *Mujer) Genero() string { return "MUJER" }
