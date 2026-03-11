# Extract Diff

A Go application that compares two text files and extracts the lines that exist in the first file but not in the second file.

## 📋 Functionality

The application performs an efficient comparison between two input files:
- **Reads two files** passed as parameters on the command line
- **Compares lines** using a map for O(1) lookups, ensuring performance
- **Extracts differences**: finds all lines that are in the first file but not in the second
- **Generates output** in a timestamped file (date and time), stored in `./results/`

## 🚀 How to Compile

```bash
# Compile the application
go build -o extract-diff main.go

# Or run directly (without compiling)
go run main.go <file1> <file2>
```

## 💻 How to Use

### Basic Syntax

```bash
./extract-diff <file1> <file2>
```

### Where:
- **file1**: first file for comparison (lines that are not in file2 will be extracted)
- **file2**: second file for comparison

Files must be located in the `./arqs/` directory

## 📂 Directory Structure

```
./
├── arqs/              # Directory for input files
├── results/           # Directory where results are saved (created automatically)
├── extract-diff       # Compiled executable
├── main.go            # Source code
└── README.md
```

## 📝 Usage Examples

### Example 1: Basic Comparison

**Your file `./arqs/file1.txt`:**
```
apple
banana
cherry
date
elderberry
```

**Your file `./arqs/file2.txt`:**
```
apple
cherry
fig
```

**Command:**
```bash
./extract-diff file1.txt file2.txt
```

**Output:**
- File created: `./results/diff__2026-03-11_14-30` (with current date/time)
- Content of the resulting file:
```
banana
date
elderberry
```

### Example 2: URL List Comparison

**Your file `./arqs/old_urls.txt`:**
```
https://example.com/page1
https://example.com/page2
https://example.com/page3
https://example.com/page4
```

**Your file `./arqs/new_urls.txt`:**
```
https://example.com/page2
https://example.com/page4
```

**Command:**
```bash
./extract-diff old_urls.txt new_urls.txt
```

**Result:** A file with the old URLs that were removed:
```
https://example.com/page1
https://example.com/page3
```

### Example 3: Deleted IDs Audit

**Your file `./arqs/ids_before.txt`:**
```
user_001
user_002
user_003
user_004
user_005
```

**Your file `./arqs/ids_after.txt`:**
```
user_002
user_004
```

**Command:**
```bash
./extract-diff ids_before.txt ids_after.txt
```

**Result:** File with the IDs that were deleted.

## ⚙️ Features

- ✅ **Fast**: uses map (hash) for O(1) comparison
- ✅ **Practical**: automatically creates necessary directories
- ✅ **Trackable**: names results with timestamp for easy organization
- ✅ **Simple**: straightforward command-line interface

## 🔧 Error Handling

The application validates:
- If exactly 2 files were passed as parameters
- If the files exist and can be read
- If directories can be created
- If the result file can be written

## 📌 Notes

- Lines are compared exactly (case-sensitive)
- Whitespace is considered in the comparison
- Empty lines are also considered
- Results are always saved with timestamp in format: `YYYY-MM-DD_HH-MM`