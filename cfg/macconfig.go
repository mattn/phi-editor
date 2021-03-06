package cfg

var DEFAULT_MAC_TOML_CONFIG string = `[editor]
tab_size = 4
hungry_backspace = true
tabs_are_spaces = true
match_braces = false
maintain_indentation = true
highlight_line = true
font_face = "Courier New Bold"
font_size = 20
show_line_numbers = true

[render]
aliased = true
accelerated = true
throttle_cpu_usage = true
always_render = true

[theme]
background = 0x002649
foreground = 0xf2f4f6
cursor = 0xf2f4f6
cursor_invert = 0x000000

[cursor]
flash_rate = 400
reset_delay = 400
draw = true
flash = true

[file_associations]
[file_associations.c]
extensions = [".c", ".h", ".cc"]

[file_associations.go]
extensions = [".go"]

[syntax.go]
[syntax.go.keyword]
colouring = 0xf0a400
match = [
	"type", "import", "package", "func", "struct",
	"append", "delete", "make", "for", "if", "while",
	"switch", "select", "chan", "else", "var", "const",
	"iota", "case"
]

[syntax.go.type]
colouring = 0xf8f273
match = [
	"int", "string", "uint",
	"int8", "int16", "int32", "int64",
	"uint8", "uint16", "uint32", "uint64",
	"rune", "byte", "float32", "float64"
]

[syntax.go.comment]
colouring = 0x4b79fc
pattern = "[\\/]+.*"

[syntax.go.symbol]
colouring = 0xf0a400
match = [
	"=", ":", ";", "*", "&", "+", "-", "/", "%",
	"^", "#", "!", "@", "<", ">", ".", ","	
]

[syntax.c]
[syntax.c.type]
colouring = 0xff0000
match = [
	"int", "char", "bool", "float", "double", "void",
	"uint8_t", "uint16_t", "uint32_t", "uint64_t",
	"int8_t", "int16_t", "int32_t", "int64_t"
]

[syntax.c.keyword]
colouring = 0xff00ff
match = [
	"for", "break", "if", "else", "continue", "return",
	"goto", "static", "extern", "const", "typedef",
]

[commands]
[commands.save]
shortcut = "super+s"

[commands.close_buffer]
shortcut = "super+w"

[commands.delete_line]
shortcut = "super+d"`
