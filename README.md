# Git in Go (Gig)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Gig es una implementación de Git escrita en Go, diseñada para aprender cómo funciona un sistema de control de versiones (VCS) y un poco de Golang.

## 🚀 Características

- Implementación de los comandos básicos de Git como:
  - Init
  - Add
  - Commit
- Alto rendimiento gracias a Go
- Fácil de extender y personalizar

## 📋 Requisitos

- Go 1.20 o superior

## 🛠️ Instalación

```bash
go install github.com/hepem/gig@latest
```

## 📖 Uso

```bash
# Inicializar un nuevo repositorio
gig init

# Agregar un archivo a trackear
gig add <archivo>

# Crear un commit en el repositorio
gig commit -m "<mensaje>"
```

## 💡 Ejemplo

```bash
# Crear un nuevo directorio para nuestro proyecto
mkdir my-app
cd my-app

# Inicializar un nuevo repositorio Gig
gig init
```

Esto crea un directorio del estilo:
```
my-app
├── .gig
│   ├── objects
│   │   ├── info
│   │   └── packs
│   ├── refs
│   │   ├── heads
│   │   └── tags
│   └── HEAD
├── src
│   └── main.py
├── requirements.txt
└── README.md
```

```bash
# Realizamos cambios en nuestro main.py
echo "print('Hello World!')" > src/main.py

# Agregar los archivos al área de preparación
gig add src/main.py
```

Esto crea un nuevo directorio en .gig/objects:
```
my-app
├── .gig
│   ├── objects
│   │   ├── ac
│   │   │   └── c43e03eca34ed127049e1849ec726df47649ec
│   │   ├── info
│   │   └── packs
│   ├── refs
│   │   ├── heads
│   │   └── tags
│   ├── HEAD
│   └── index
├── src
│   └── main.py
├── requirements.txt
└── README.md
```

Creando un nuevo archivo .gig/index con el siguiente contenido:
```
acc43e03eca34ed127049e1849ec726df47649ec src/main.py
```

```bash
# Crear un commit con los cambios
gig commit -m "Primer commit: archivos iniciales"
```

Esto crea un nuevo commit en el repositorio, dejando una estructura:
```
my-app
├── .gig
│   ├── objects
│   │   ├── 14
│   │   │   └── 33d7e5f08871340e6fd3147525765de4cbfe4a
│   │   ├── 67
│   │   │   └── b59b4d419cfba7cd0b971ce46858c18b0eed1c
│   │   ├── ac
│   │   │   └── c43e03eca34ed127049e1849ec726df47649ec
│   │   ├── info
│   │   └── packs
│   ├── refs
│   │   ├── heads
│   │   │   └── master
│   │   └── tags
│   ├── HEAD
│   └── index
├── src
│   └── main.py
├── requirements.txt
└── README.md
```

El archivo index queda completamente vacío y el archivo .gig/refs/heads/master con el siguiente contenido:

```
67b59b4d419cfba7cd0b971ce46858c18b0eed1c
```


## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## 🙏 Recursos utilizados

- [Build your own X](https://github.com/codecrafters-io/build-your-own-x?tab=readme-ov-file#build-your-own-git)
- [Rebuilding Git in Ruby](https://thoughtbot.com/blog/rebuilding-git-in-ruby)
