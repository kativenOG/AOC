{
	description="Flake for Advent of Code in golang!";

  	inputs={
  		nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  		flake-utils.url = "github:numtide/flake-utils";
  	};
  	outputs = {self, nixpkgs, flake-utils, ...}@inputs:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
	  {
		devShells.default = pkgs.mkShell {
          nativeBuildInputs = [
		  	pkgs.go 
			pkgs.gore
			pkgs.gopls
          ];
		  shellhook=''
		  	gofmt -w . 
		  '';
		};
      }
	);
}





