
# Note to users
I built this mostly as practice for some other projects before realizing there were a lot of competing options under the same name. I won't be maintaining this and you should really use a different one anyhow. 

# golings 🐹
Small exercises to get you used to reading and writing Go code - Recommended in parallel to reading [the official Go tour](https://go.dev/tour) and [Effective Go](https://go.dev/doc/effective_go)

## Getting Started

### Prerequisites
- Basic understanding of programming concepts
- A terminal or command prompt
- Internet access for documentation

### Getting Started

#### Installation

1. **Install Go**
   Follow the official installation guide: https://go.dev/doc/install

2. **Add Go bin directory to PATH**
   Add this to your shell configuration (.bashrc, .zshrc, etc.):
   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```
   Then reload your shell:
   ```bash
   source ~/.bashrc  # or ~/.zshrc
   ```

3. **Clone this repository**
   ```bash
   git clone https://github.com/[yourusername]/golings
   cd golings
   ```

4. **Install golings CLI**
   ```bash
   go install ./cmd/golings
   ```

#### Using the Progress Tracker

1. **List all exercises**
   ```bash
   golings list
   ```

2. **Verify an exercise**
   ```bash
   golings verify <exercise_name>
   ```

3. **Show your progress**
   ```bash
   golings progress
   ```

4. **Run a specific exercise**
   ```bash
   golings run <exercise_name>
   ```

The progress tracker will show:
- Current completion status
- Recently completed exercises
- Overall progress percentage

### Recommended Learning Path

We recommend going through the exercises in order, starting with the basics and progressing through more advanced topics. Each exercise includes a link to relevant Go documentation.

### Contributing

Contributions are welcome! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.
