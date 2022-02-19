package available

import "entgo.io/ent/schema"

type Annotation struct {
	Activation   bool
	Lifespan     bool
	SoftDeletion bool
}

func (Annotation) Name() string {
	return "Available"
}

func (a Annotation) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}

	if ant.Activation {
		a.Activation = true
	}
	if ant.Lifespan {
		a.Lifespan = true
	}
	if ant.SoftDeletion {
		a.SoftDeletion = true
	}

	return a
}
