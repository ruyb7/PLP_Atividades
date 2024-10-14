package main

type Professor struct {
    nome     string
    escolas  []*Escola
}

type Escola struct {
    nome       string
    professores []*Professor
}

func (p *Professor) AdicionarEscola(e *Escola) {
    p.escolas = append(p.escolas, e)
    e.AdicionarProfessor(p)
}

func (e *Escola) AdicionarProfessor(p *Professor) {
    for _, professor := range e.professores {
        if professor == p {
            return
        }
    }
    e.professores = append(e.professores, p)
}

func main() {}
