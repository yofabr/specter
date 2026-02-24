# Specter

Specter is a concurrent port discovery engine designed to map the network surface of a local or remote host with minimal latency and high resource efficiency.

## Project Structure

---------------------------------------------------------------------
  
   ██████╗ ███████╗███████╗██╗     ██╗     ███████╗██╗  ██╗
  ██╔════╝ ██╔════╝██╔════╝██║     ██║     ██╔════╝╚██╗██╔╝
  ██║  ███╗█████╗  █████╗  ██║     ██║     █████╗   ╚███╔╝ 
  ██║   ██║██╔══╝  ██╔══╝  ██║     ██║     ██╔══╝   ██╔██╗ 
  ╚██████╔╝███████╗███████╗███████╗███████╗███████╗██╔╝ ██╗
   ╚═════╝ ╚══════╝╚══════╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝
  
  ---------------------------------------------------------------------

	Tiny, Fast, and Deployable anywhere — automate the mundane, unleash your creativity

```
specter/
├── cmd/
│   └── specter/
│       └── main.go          # CLI entry point
├── internal/
│   └── scanner/
│       ├── scanner.go       # Core scanning logic
│       └── scanner_test.go  # Unit tests
├── go.mod
├── Makefile
└── README.md
```

## Usage

Here are the steps to use Specter:

1.  **Build the application:**

    Run the following command to build the application:

    ```bash
    make build
    ```

    This will create an executable file named `specter` (or `specter.exe` on Windows) in the current directory.

2.  **Run the application:**

    To scan the ports on the default target IP address (127.0.0.1), run the following command:

    ```bash
    ./specter
    ```

3.  **CLI Options:**

    | Flag      | Default     | Description                    |
    |-----------|-------------|--------------------------------|
    | `-target` | 127.0.0.1   | Target IP address              |
    | `-start`  | 1           | Start port                     |
    | `-end`    | 1024        | End port                       |
    | `-workers`| 100         | Number of concurrent workers   |

    Example:

    ```bash
    ./specter -target 192.168.1.1 -start 1 -end 65535 -workers 200
    ```

## Installation

You can install Specter by following these steps:

1.  **Build the application:**

    Run the following command to build the application:

    ```bash
    make build
    ```

2.  **Install the application:**

    Run the following command to install the application:

    ```bash
    sudo make install
    ```

    This will build the application and move it to `/usr/local/bin`, making it available system-wide.

## License

MIT