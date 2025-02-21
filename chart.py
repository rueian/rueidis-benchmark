#!/usr/bin/env python3
import sys, re
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
import numpy as np
from io import BytesIO

def extract_filenames(content):
    """
    Look for a header line containing '.txt' and extract all tokens that look like file names.
    """
    for line in content.splitlines():
        if ".txt" in line:
            parts = line.split("│")
            files = [p.strip() for p in parts if ".txt" in p]
            if files:
                return files
    return []

def parse_section(lines, metric, n_files):
    """
    For each data line, extract the test name (the first token) and the first n_files metric values.
    For sec/op, values in nanoseconds are converted to microseconds (1µs = 1000n).
    """
    test_names = []
    data = []
    # Capture a numeric value with its unit before the "±" symbol.
    pattern = re.compile(r'([0-9.]+)([a-zµ]*)\s*±')
    for line in lines:
        line = line.strip()
        if not line or line.startswith("geomean"):
            continue
        parts = line.split()
        test_name = parts[0]
        matches = pattern.findall(line)
        if len(matches) < n_files:
            continue
        values = matches[:n_files]
        if metric == "sec/op":
            conv = []
            for val, unit in values:
                f = float(val)
                if unit == "µ":
                    conv.append(f)
                elif unit == "n":
                    conv.append(f / 1000.0)
                else:
                    conv.append(f)
            conv_values = conv
        else:
            conv_values = [float(v) for v, _ in values]
        test_names.append(test_name)
        data.append(conv_values)
    return test_names, data

def main():
    content = sys.stdin.read()
    files = extract_filenames(content)
    if not files:
        sys.stderr.write("Warning: No file names found in input. Using default names.\n")
        files = ["File1", "File2", "File3"]
    n_files = len(files)

    sections = re.split(r'\n\s*\n', content)
    data_sections = {}
    for section in sections:
        if "sec/op" in section:
            metric = "sec/op"
        elif "B/op" in section:
            metric = "B/op"
        elif "allocs/op" in section:
            metric = "allocs/op"
        else:
            continue
        lines = section.splitlines()
        data_lines = []
        for line in lines:
            if line.startswith(("goos:", "goarch:", "pkg:", "cpu:")):
                continue
            if line.strip().startswith("│"):
                continue
            if "±" not in line:
                continue
            data_lines.append(line)
        test_names, dataset = parse_section(data_lines, metric, n_files)
        if test_names:
            data_sections[metric] = (test_names, dataset)

    # Create a figure with one subplot per metric (wider figure).
    fig, axes = plt.subplots(nrows=3, ncols=1, figsize=(20, 15))
    metrics = ["sec/op", "B/op", "allocs/op"]

    for ax, metric in zip(axes, metrics):
        if metric not in data_sections:
            ax.set_visible(False)
            continue
        test_names, dataset = data_sections[metric]
        dataset = np.array(dataset)
        n_tests = len(test_names)
        n_bars = dataset.shape[1]
        indices = np.arange(n_tests)
        bar_height = 0.8 / n_bars

        for i in range(n_bars):
            ax.barh(indices - 0.4 + i * bar_height + bar_height/2,
                    dataset[:, i],
                    height=bar_height,
                    label=files[i] if i < len(files) else f"File{i+1}")
        ax.set_yticks(indices)
        ax.set_yticklabels(test_names)
        ax.invert_yaxis()
        ax.set_title(metric)
        ax.legend()
        ax.set_xlabel("Time (µs)" if metric == "sec/op" else metric)
        # Format tick labels with two decimal places.
        ax.xaxis.set_major_formatter(ticker.FormatStrFormatter('%.2f'))
        # Triple the number of x-axis ticks.
        default_ticks = ax.get_xticks()
        new_tick_count = len(default_ticks) * 3 if len(default_ticks) > 0 else 10
        ax.xaxis.set_major_locator(ticker.LinearLocator(numticks=new_tick_count))
    plt.tight_layout()

    buf = BytesIO()
    plt.savefig(buf, format='png')
    buf.seek(0)
    sys.stdout.buffer.write(buf.getvalue())

if __name__ == "__main__":
    main()
