# Specter

Specter is a concurrent port discovery engine designed to map the network surface of a local or remote host with minimal latency and high resource efficiency.

## Usage

To use Specter, you first need to build the application. You can do this by running the command `make build`. This will create an executable file named `specter` (or `specter.exe` on Windows).

To run the application, you can use the command `./specter`. This will scan the ports on the default target IP address, which is `127.0.0.1`.

To specify a different target IP address, you can use the `-target` flag. For example, to scan the ports on `192.168.1.1`, you would use the command `./specter -target 192.168.1.1`.

### Installation

You can install Specter by running the command `make install`. This will build the application and move it to `/usr/local/bin`, making it available system-wide.

## License

MIT
