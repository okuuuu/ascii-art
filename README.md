<p style='text-align: justify;'>

# ASCII ART

The current ASCII art project includes following subprojects: 
-   [x] reverse
-   [x] color
-   [x] output
-   [x] fs
-   [x] justify

res folder contains some functions required for the main code functioning, additional files and structures. test folder contains the txt test files for audit.

The programm is launched from the main.go file.

## Test

Run `bash test.sh` while in the root folder.

## Reverse

While **--reverse** flag is active, the other flags cannot be used. You can still generate rune map using different banners.
Compatible with: banners.

## Color

Color flags can take or not to take the following argument as a fragment from the string to be colored. Default color is white. If multiple colors were applied on the same fragment, the last declared color will be shown in the output For example, the allowed color flag combinations are:

-   `--color=green` ---> displays green color text
-   `--color=green christ` ---> displays white color text with every christ fragment colored in green
-   `--color=green --color=red christ` --color=yellow mas ---> displays green color text with every
-   `--color=green --color=red` ---> displays red color text
    christ fragment colored in red and every mas fragment colored in yellow

List of all allowed colors is in res/colors.go file.
Compatible with: banners, --align.

## Output

Color flags do not work when the output is a file. However different banner selection is still available.
Compatible with: banners.

## Banners

You can choose every existing banner by specifying its name in the last command line argument (without txt). Please note, that choosing the string for ASCII representation has higher order of precedence. That means if arguments `--color=red sha shadow` were passed then the output will be ASCII art of the word "shadow" where "sha" is red color.
Compatible with: --reverse, --color, --output, --align.

## Justify

Insert the --align flag, to align the terminal output left, right, center or justify it. If the result string will not fit the terminal window, the efect of the flag will be neglected. 
Compatible with: --color, banners.
</p>

@okuu