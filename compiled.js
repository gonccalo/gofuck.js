var mem 	= new Array(30000).fill(0);
var pointer = 0;
var outDiv  = document.getElementById("output");
var enter   = document.getElementById("enter");
var input   = document.getElementById("input");
var info    = document.getElementById("info");
var gen 	= main();
gen.next();
enter.onclick = function() {
	mem[pointer] = input.value.charCodeAt(0);
	gen.next();
};
function* main() {
	pointer++;
	pointer++;
	pointer++;
	mem[pointer]++;
	while (mem[pointer] != 0) {
		while (mem[pointer] != 0) {
			mem[pointer]--;
		}
		pointer++;
		pointer++;
		while (mem[pointer] != 0) {
			mem[pointer]--;
		}
		mem[pointer]++;
		mem[pointer]++;
		pointer++;
		mem[pointer]++;
		pointer++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		while (mem[pointer] != 0) {
			pointer--;
			mem[pointer]++;
			mem[pointer]++;
			mem[pointer]++;
			mem[pointer]++;
			pointer++;
			pointer++;
			mem[pointer]++;
			mem[pointer]++;
			pointer--;
			mem[pointer]--;
		}
		mem[pointer]++;
		mem[pointer]++;
		pointer++;
		pointer++;
		mem[pointer]++;
		pointer++;
		mem[pointer]++;
		pointer++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		mem[pointer]++;
		while (mem[pointer] != 0) {
			pointer++;
			mem[pointer]++;
			mem[pointer]++;
			pointer++;
			mem[pointer]++;
			mem[pointer]++;
			mem[pointer]++;
			mem[pointer]++;
			mem[pointer]++;
			mem[pointer]++;
			pointer--;
			pointer--;
			mem[pointer]--;
		}
		mem[pointer]++;
		pointer++;
		pointer++;
		pointer++;
		enter.disabled = false;
		info.innerText = "Input text now";
		yield;
		info.innerText = "";
		enter.disabled = true;
		pointer--;
		mem[pointer]++;
		mem[pointer]++;
		while (mem[pointer] != 0) {
			while (mem[pointer] != 0) {
				pointer++;
				while (mem[pointer] != 0) {
					mem[pointer]--;
					pointer++;
					pointer++;
				}
				pointer--;
				while (mem[pointer] != 0) {
					pointer++;
					pointer++;
				}
				pointer--;
				pointer--;
				mem[pointer]--;
			}
			pointer--;
			while (mem[pointer] != 0) {
				pointer--;
			}
			pointer--;
			mem[pointer]++;
			pointer++;
			pointer++;
			while (mem[pointer] != 0) {
				pointer++;
			}
			pointer++;
			while (mem[pointer] != 0) {
				pointer--;
				mem[pointer]++;
				pointer++;
				mem[pointer]--;
				while (mem[pointer] != 0) {
					while (mem[pointer] != 0) {
						pointer--;
						mem[pointer]++;
						pointer++;
						mem[pointer]--;
					}
					pointer++;
				}
				pointer--;
				while (mem[pointer] != 0) {
					while (mem[pointer] != 0) {
						while (mem[pointer] != 0) {
							mem[pointer]--;
						}
						pointer--;
					}
					mem[pointer]++;
					mem[pointer]++;
					pointer--;
					mem[pointer]--;
					while (mem[pointer] != 0) {
						pointer--;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						mem[pointer]++;
						pointer++;
						while (mem[pointer] != 0) {
							pointer--;
							mem[pointer]--;
							pointer++;
							mem[pointer]--;
						}
						pointer++;
						pointer++;
					}
					pointer++;
					pointer++;
				}
			}
			pointer--;
			pointer--;
		}
		pointer--;
	}
	pointer--;
	while (mem[pointer] != 0) {
		while (mem[pointer] != 0) {
			pointer--;
		}
		pointer++;
		while (mem[pointer] != 0) {
			while (mem[pointer] != 0) {
				pointer++;
			}
			pointer++;
			pointer++;
			while (mem[pointer] != 0) {
				pointer++;
				pointer++;
			}
			mem[pointer]++;
			while (mem[pointer] != 0) {
				pointer--;
				pointer--;
			}
			pointer--;
			while (mem[pointer] != 0) {
				pointer--;
			}
			pointer--;
			mem[pointer]++;
			pointer++;
			pointer++;
			mem[pointer]--;
		}
		pointer++;
		while (mem[pointer] != 0) {
			pointer++;
		}
		mem[pointer]++;
		while (mem[pointer] != 0) {
			mem[pointer]--;
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			while (mem[pointer] != 0) {
				pointer--;
				pointer--;
			}
			pointer--;
			while (mem[pointer] != 0) {
				pointer--;
			}
			mem[pointer]++;
			pointer--;
			pointer--;
			while (mem[pointer] != 0) {
				mem[pointer]++;
				pointer++;
				mem[pointer]++;
				pointer--;
				pointer--;
				mem[pointer]--;
				while (mem[pointer] != 0) {
					pointer++;
					mem[pointer]--;
					mem[pointer]--;
					pointer++;
					mem[pointer]++;
					pointer--;
					pointer--;
					mem[pointer]--;
					while (mem[pointer] != 0) {
						pointer++;
						mem[pointer]++;
						pointer--;
						while (mem[pointer] != 0) {
							pointer++;
							pointer++;
							mem[pointer]++;
							pointer--;
							pointer--;
							mem[pointer]--;
						}
					}
				}
				pointer++;
				while (mem[pointer] != 0) {
					pointer--;
					mem[pointer]++;
					pointer++;
					mem[pointer]--;
				}
				pointer--;
			}
			mem[pointer]++;
			mem[pointer]++;
			pointer++;
			pointer++;
			mem[pointer]--;
			mem[pointer]--;
			pointer++;
			while (mem[pointer] != 0) {
				pointer++;
			}
			pointer++;
			pointer++;
			while (mem[pointer] != 0) {
				pointer++;
				pointer++;
			}
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			pointer++;
			pointer++;
			mem[pointer]++;
			pointer--;
			while (mem[pointer] != 0) {
				while (mem[pointer] != 0) {
					pointer--;
				}
				pointer--;
			}
			pointer++;
			while (mem[pointer] != 0) {
				while (mem[pointer] != 0) {
					pointer--;
					pointer--;
				}
				pointer--;
				while (mem[pointer] != 0) {
					pointer--;
				}
				mem[pointer]++;
				while (mem[pointer] != 0) {
					mem[pointer]--;
					pointer--;
					mem[pointer]++;
					pointer++;
					pointer++;
					mem[pointer]--;
					while (mem[pointer] != 0) {
						pointer--;
						pointer--;
						mem[pointer]++;
						pointer++;
						mem[pointer]++;
						mem[pointer]++;
						pointer++;
						mem[pointer]--;
						while (mem[pointer] != 0) {
							pointer--;
							mem[pointer]--;
							pointer++;
							while (mem[pointer] != 0) {
								pointer--;
								pointer--;
								mem[pointer]++;
								pointer++;
								pointer++;
								mem[pointer]--;
							}
						}
					}
					pointer--;
					while (mem[pointer] != 0) {
						pointer++;
						mem[pointer]++;
						pointer--;
						mem[pointer]--;
					}
					pointer++;
				}
				pointer++;
				while (mem[pointer] != 0) {
					pointer++;
				}
				pointer++;
			}
			pointer++;
			while (mem[pointer] != 0) {
				pointer++;
				pointer++;
			}
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			pointer++;
			pointer++;
			mem[pointer]++;
			pointer++;
			pointer++;
			mem[pointer]++;
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			mem[pointer]--;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			pointer++;
			outDiv.innerText = outDiv.innerText + String.fromCharCode(mem[pointer]);
			console.log(String.fromCharCode(mem[pointer]));
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			pointer++;
			mem[pointer]--;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			pointer++;
			enter.disabled = false;
			info.innerText = "Input text now";
			yield;
			info.innerText = "";
			enter.disabled = true;
			pointer++;
			pointer++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			pointer++;
			mem[pointer]++;
			pointer++;
		}
		pointer--;
		pointer--;
		while (mem[pointer] != 0) {
			mem[pointer]++;
			pointer--;
			pointer--;
		}
		pointer--;
	}
}
