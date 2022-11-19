let systemPkgs = import <nixpkgs> {
	overlays = [
		(self: super: {
			go = super.go_1_19;
		})
	];
};

in { pkgs ? systemPkgs }:

let fetch-samples = pkgs.writeShellScriptBin "fetch-samples" ''
		set -e
		if [[ "$1" == "" ]]; then
			echo "Usage: fetch-samples <url>"
		fi

		wget -O samples.zip "$1"
		unzip samples.zip
		rm samples.zip
	'';

	mainTemplate = pkgs.writeText "mainTemplate.go" ''package main

func main() {
	ncases, _ := strconv.Atoi(readLine())
	for n := 0; n < ncases; n++ {
	}
}

var stdinBuf = bufio.NewReader(os.Stdin)

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}
	'';

	new-challenge = pkgs.writeShellScriptBin "new-challenge" ''
		set -e
		if [[ "$1" == "" ]]; then
			echo "Usage: new-challenge <name>"
		fi

		mkdir -p "$1"
		cd "$1"
		cp "${mainTemplate}" main.go
	'';

in pkgs.mkShell {
	buildInputs = with pkgs; [
		go
		gopls
		gotools
		unzip
		wget

		fetch-samples
		new-challenge
	];

	DEBUG = "1";
}
