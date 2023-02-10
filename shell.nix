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

	# clangd hack.
	llvmPackages = pkgs.llvmPackages_latest;
	clang-unwrapped = llvmPackages.clang-unwrapped;
	clang  = llvmPackages.clang;
	clangd = pkgs.writeScriptBin "clangd" ''
	    #!${pkgs.stdenv.shell}
		export CPATH="$(${clang}/bin/clang -E - -v <<< "" \
			|& ${pkgs.gnugrep}/bin/grep '^ /nix' \
			|  ${pkgs.gawk}/bin/awk 'BEGIN{ORS=":"}{print substr($0, 2)}' \
			|  ${pkgs.gnused}/bin/sed 's/:$//')"
		export CPLUS_INCLUDE_PATH="$(${clang}/bin/clang++ -E - -v <<< "" \
			|& ${pkgs.gnugrep}/bin/grep '^ /nix' \
			|  ${pkgs.gawk}/bin/awk 'BEGIN{ORS=":"}{print substr($0, 2)}' \
			|  ${pkgs.gnused}/bin/sed 's/:$//')"
	    ${clang-unwrapped}/bin/clangd
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

		clangd
		clang

		python3
		pyright
	];

	DEBUG = "1";
}
