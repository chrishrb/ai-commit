{
  description = "ai-commit - generate your commit messages with the help of ai";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, ... }@inputs: inputs.flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        ai-commit = pkgs.buildGoModule {
          name = "ai-commit";
          src = self;
          vendorHash = "sha256-mr5Q2iYL56B8VC/gRi3rP//PdD3aFZc+5obn63QkWWk=";
        };
      in
      {
        packages.default = ai-commit;
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = [ ai-commit ];
        };
      }
    );
}
