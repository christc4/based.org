# Picom

## Transparency

These are my personal configs 

	# Enable/disable Picom
	backend = "xrender";

	# General settings
	vsync = true;
	backend = "glx";
	glx-no-stencil = true;
	use-damage = true;


	# Shadow settings
	shadow = false;

	# Transparency for inactive windows
	inactive-opacity = 0.8;
	active-opacity = 0.95;

	# Other optimizations
	unredir-if-possible = true;

	opacity-rule = [
	  "100:name *= 'mpv'"
	];


## Greyscale

Create `file.gsl`

	#version 330

	in vec2 texcoord;
	uniform sampler2D tex;
	uniform float opacity;

	vec4 default_post_processing(vec4 c);

	vec4 window_shader() {
		vec2 texsize = textureSize(tex, 0);
		vec4 color = texture2D(tex, texcoord / texsize, 0);

		color = vec4(vec3(0.2126 * color.r + 0.7152 * color.g + 0.0722 * color.b) * opacity, color.a * opacity);

		return default_post_processing(color);
	}

Start picom by including the .gsl file path

`picom --backend glx --window-shader-fg /path/to/shader/file.gls`


