package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// Estrutura para representar uma música
type Musica struct {
	Titulo     string
	ArquivoMP3 string
}

// Estrutura para representar um nó na lista circular
type No struct {
	musica  Musica
	proximo *No
}

// Estrutura para representar a lista circular
type ListaCircular struct {
	cabeca *No
}

// Função para adicionar uma música à lista circular
func (lista *ListaCircular) Adicionar(musica Musica) {
	no := &No{musica: musica}
	if lista.cabeca == nil {
		lista.cabeca = no
		no.proximo = lista.cabeca
	} else {
		atual := lista.cabeca
		for atual.proximo != lista.cabeca {
			atual = atual.proximo
		}
		atual.proximo = no
		no.proximo = lista.cabeca
	}
}

// Função para imprimir a lista circular de músicas
func (lista *ListaCircular) Imprimir() {
	if lista.cabeca == nil {
		return
	}
	atual := lista.cabeca
	for {
		fmt.Println(atual.musica.Titulo)
		atual = atual.proximo
		if atual == lista.cabeca {
			break
		}
	}
}

// Função para remover uma música da lista circular pelo título
func (lista *ListaCircular) Remover(titulo string) {
	if lista.cabeca == nil {
		return
	}
	atual := lista.cabeca
	anterior := lista.cabeca

	// Encontrar o nó com a música a ser removida
	for atual.musica.Titulo != titulo {
		anterior = atual
		atual = atual.proximo
		if atual == lista.cabeca {
			// Caso a música não seja encontrada na lista
			return
		}
	}

	if atual == lista.cabeca {
		// Remover a cabeça da lista
		proximoNo := lista.cabeca.proximo
		for anterior.proximo != lista.cabeca {
			anterior = anterior.proximo
		}
		lista.cabeca = proximoNo
		anterior.proximo = lista.cabeca
	} else {
		anterior.proximo = atual.proximo
	}
}

// Função para tocar uma música
func tocarMusica(arquivoMP3 string) error {
	// Abre o arquivo de áudio
	f, err := os.Open(arquivoMP3)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo %s: %v", arquivoMP3, err)
	}
	defer f.Close()

	// Decodifica o arquivo mp3
	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		return fmt.Errorf("erro ao decodificar o arquivo mp3 %s: %v", arquivoMP3, err)
	}

	// Inicializa o contexto de áudio
	context, err := oto.NewContext(decoder.SampleRate(), 2, 2, 4096)
	if err != nil {
		return fmt.Errorf("erro ao criar o contexto de áudio: %v", err)
	}
	defer context.Close()

	// Cria um reprodutor
	player := context.NewPlayer()
	defer player.Close()

	// Escreve o áudio decodificado no reprodutor
	if _, err := io.Copy(player, decoder); err != nil {
		return fmt.Errorf("erro ao reproduzir o áudio: %v", err)
	}

	return nil
}

// Função para tocar as músicas da lista circular
func (lista *ListaCircular) TocarTodas(repeticoes int) {
	if lista.cabeca == nil {
		return
	}
	atual := lista.cabeca
	for i := 0; i < repeticoes; i++ {
		for {
			fmt.Println("Tocando música:", atual.musica.Titulo)
			if err := tocarMusica(atual.musica.ArquivoMP3); err != nil {
				fmt.Println("Erro ao tocar música:", err)
			}
			atual = atual.proximo
			if atual == lista.cabeca {
				break
			}
		}
	}
}

func main() {
	listaCircular := ListaCircular{}

	// Adicionar músicas à lista circular
	listaCircular.Adicionar(Musica{Titulo: "Dorothy - Black Sheep", ArquivoMP3: "C:\\Users\\maria\\Downloads\\Dorothy - Black Sheep (320).mp3"})
	listaCircular.Adicionar(Musica{Titulo: "Anjulie - Boom", ArquivoMP3: "C:\\Users\\maria\\Downloads\\Anjulie - Boom (lyrics) (320).mp3"})
	listaCircular.Adicionar(Musica{Titulo: "Serena Ryder - Got Your Number", ArquivoMP3: "C:\\Users\\maria\\Downloads\\Serena Ryder - Got Your Number (Official Video) (320).mp3"})
	listaCircular.Adicionar(Musica{Titulo: "Nova Twins - Antagonist", ArquivoMP3: "C:\\Users\\maria\\Downloads\\Nova Twins - Antagonist (Official Audio) (320).mp3"})
	listaCircular.Adicionar(Musica{Titulo: "Black Sheep (Brie Larson Vocal Version)", ArquivoMP3: "C:\\Users\\maria\\Downloads\\Black Sheep (Brie Larson Vocal Version) (320).mp3"})

	// Imprimir a lista de músicas
	fmt.Println("Lista de músicas:")
	listaCircular.Imprimir()

	// Tocar todas as músicas da lista circular 2 vezes
	fmt.Println("\nTocando todas as músicas vezes:")
	listaCircular.TocarTodas(2)

	// Remover uma música da lista
	listaCircular.Remover("Nova Twins - Antagonist")
	fmt.Println("\nLista de músicas após remover 'Nova Twins - Antagonist':")
	listaCircular.Imprimir()

	// Tocar todas as músicas da lista circular novamente
	fmt.Println("\nTocando todas as músicas novamente:")
	listaCircular.TocarTodas(1)
}
