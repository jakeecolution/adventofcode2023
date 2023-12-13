import re

repl_trans = {
    'one': '1e',
    'two': '2o',
    'three': '3e',
    'four': '4r',
    'five': '5e',
    'six': '6x',
    'seven': '7n',
    'eight': '8t',
    'nine': '9e',
}
translate = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9',
}

ok_str = ('1', '2', '3', '4', '5', '6', '7', '8', '9')
regex = r"([1-9]|one|two|three|four|five|six|seven|eight|nine)"

def read_line(line):
    total = ""
    pattern = re.compile(regex)
    result = pattern.findall(line)
    # print(f"before: '{line}'")
    if result[0] in ok_str:
        total = result[0]
    else:
        replace = repl_trans[result[0]]
        total = translate[result[0]]
        line = re.sub(regex, replace, line, count=1)
        # print(f"after: '{line}'")
    result = pattern.findall(line)
    if result[-1] in ok_str:
        total += result[-1]
    else:
        total += translate[result[-1]]
    try:
        out = int(total)
    except ValueError:
        out = 0
    # print(f"total: {total}, {out}")
    return out


def main():
    sum = 0
    with open('analyzeme.txt') as f:
        for line in f:
            line = line.strip()
            sum += read_line(line)
    print(sum)
    sum = 0
    with open('input.txt', 'r') as f:
        for line in f:
            line = line.strip()
            sum += read_line(line)
    print(sum)

if __name__ == '__main__':
    main()