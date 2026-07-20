package catalogo

// En Go no hay herencia: la jerarquía se expresa con una interface...
type Product interface {
	ProductID() int64
	Describe() string
}

// ...y structs concretos que comparten lo común por composición.
type baseProduct struct {
	ID         int64
	Name       string
	PriceCents int64
}

func (b baseProduct) ProductID() int64 { return b.ID }

type PhysicalProduct struct {
	baseProduct
	WeightGrams int
	Stock       int
}

func (p PhysicalProduct) Describe() string {
	return p.Name + " (físico, envío disponible)"
}

type DigitalProduct struct {
	baseProduct
	DownloadURL string
	FileBytes   int64
}

func (d DigitalProduct) Describe() string {
	return d.Name + " (digital, descarga inmediata)"
}
