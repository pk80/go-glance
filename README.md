# go-glance

## One Resource
### Basic
- How to think like a Go developer
- Native data types
- Native data structures
- Functions and methods
- Object Oriented Programming(?)
- Error handling
- Interfaces
- Vendoring
- Best practices/ How to do code reviews
- Tools:
  - Build environment
  - CI/CD options
  - Debuggers
  - Performance tools
  - Static analysis
  - Race detection
  - API docs

### Advanced
- Net HTTP (native)
- REST APIs with Gorilla
- JSON/YAML handling
- Logging
- gRPC
- Concurrency (goroutines and channels)
- Testing (native)
- Ginkgo (overview)
- Debugging
- Some Kubernetes code reading to see how production code is written


## Two Resource 
1. OVERVIEW
    a. Features of Go Programming
    b. Features Excluded Intentionally
    c. Go Programs.
    d. Compiling and Executing Go Programs
2. ENVIRONMENT SETUP
    a. Local Environment Setup
    b. Text Editor
    c. The Go Compiler
    d. Download Go Archive
    e. Installation on UNIX/Linux/Mac OS X, and FreeBSD
    f. Installation on Windows
    g. Verifying the Installation
3. PROGRAM STRUCTURE
    a. Hello World Example
    b. Executing a Go Program
4. BASIC SYNTAX
    a. Tokens in Go
    b. Line Separator
    c. Comments
    d. Identifiers
    e. Keywords
    f. Whitespace in Go
5. DATA TYPES
    a. Integer Types
    b. Floating Types
    c. Other Numeric Types
6. VARIABLES
    a. Variable Definition in Go
    b. Static Type Declaration in Go
    c. Dynamic Type Declaration / Type Inference in Go
    d. Mixed Variable Declaration in Go
    e. The lvalues and the rvalues in Go
7. CONSTANTS
    a. Integer Literals
    b. Floating-point Literals
    c. Escape Sequence
    d. String Literals in Go
    e. The const Keyword
8. OPERATORS
    a. Arithmetic Operators
    b. Relational Operators
    c. Logical Operators
    d. Bitwise Operators
    e. Assignment Operators
    f. Miscellaneous Operators
    g. Operators Precedence in Go
9. DECISION MAKING
    a. The if Statement
    b. The if…else Statement
    c. Nested if Statement
    d. The Switch Statement
    e. The Select Statement
    f. The if...else if...else Statement
10. LOOPS
    a. for Loop
    b. Nested for Loops
    c. Loop Control Statements
    d. The continue Statement
    e. The goto Statement.
    f. The Infinite Loop
11. FUNCTIONS
    a. Defining a Function
    b. Calling a Function
    c. Returning Multiple Values from Function
    d. Function Arguments
    e. Call by Value
    f. Call by Reference
    g. Function Usage
    h. Function Closures
    i. Method
12. SCOPE RULES
    a. Local Variables
    b. Global Variables
    c. Formal Parameters
    d. Initializing Local and Global Variables
13. STRINGS
    a. Creating Strings
    b. String Length
    c. Concatenating Strings
14. ARRAYS
    a. Declaring Arrays
    b. Initializing Arrays
    c. Accessing Array Elements
    d. Go Arrays in Detail
    e. Multidimensional Arrays in Go
    f. Two-Dimensional Arrays
    g. Initializing Two-Dimensional Arrays
    h. Accessing Two-Dimensional Array Elements
    i. Passing Arrays to Functions
15. POINTERS
    a. What Are Pointers?
    b. How to Use Pointers?
    c. Nil Pointers in Go
    d. Go Pointers in Detail
    e. Go – Array of Pointers
    f. Go – Pointer to Pointer
    g. Go – Passing Pointers to Functions
16. STRUCTURES
    a. Defining a Structure
    b. Accessing Structure Members
    c. Structures as Function Arguments
    d. Pointers to Structures
17. SLICES
    a. Defining a slice
    b. len() and cap() functions
    c. Nil slice
    d. Subslicing
    e. append() and copy() Functions
18. RANGE
19. MAPS
    a. Defining a Map
    b. delete() Function
20. RECURSION
21. TYPE CASTING
22. INTERFACES
23. ERROR HANDLING
24. Packages
25. Concurrency – Goroutines & Channel

## Three Resource
### Getting started
1. Section 1 - Getting Started
   - Go Runtime and Compilations
   - Keywords and Identifiers
   - Constants and Variables
   - Operators and Expressions
   - Local Assignments
   - Booleans, Numerics, Characters
   - Pointers and Addresses
   - Strings
2. Section 2 - Constructs
   - if-else and switch
   - for Statements
   - Counter-controlled Iterations
   - Condition-controlled Iterations
   - Range Loops
   - Using break and continue
3. Section 3 - Functions
   - Parameters and Return Values
   - Call by Value and Reference
   - Named Return Variables
   - Blank Identifiers
   - Variable Argument Parameters
   - Using defer statements
   - Recursive Functions
   - Functions as Parameters
   - Closures
4. Section 4 - Working with Data
   - Array Literals
   - Multidimensional Arrays
   - Array Parameters
   - Slices and Slice Parameters
   - Multidimensional Slices
   - Reslicing
   - Maps and Map Parameters
   - Map Slices
   - Structures and Structure Parameters
   - Structure Tags and Fields
   - Embedded Structures
   - Recursive Structures
5. Section 5 - Methods and Interfaces
   - Method Declarations
   - Functions vs. Methods
   - Pointer and Value Receivers
   - Method Values and Expressions
   - Interface Types and Values
   - Type Assertions and Type Switches
   - Method Sets with Interfaces
   - Embedded Interfaces
   - Empty Interfaces
   - Working with Interfaces
6. Section 6 - Goroutines and Channels
   - Concurrency vs. Parallelism
   - Goroutine Functions and Lambdas
   - Wait Groups
   - Channels
   - Sending and Receiving
   - Unbuffered and Buffered Channels
   - Directional Channels
   - Multiplexing with select
   - Timers and Tickers
7. Section 7 - Packages and Testing
   - Packages and Workspaces
   - Exporting Package Names
   - Import Paths and Named Imports
   - Package Initializations
   - Blank Imports
   - Unit Testing with Test Functions
   - Table Tests and Random Tests
   - Benchmarking
8. Section 8 - Working with Go
   - Files and Directories
   - Reading Directories
   - Reading Files
   - Writing Files
   - Copying Files
   - Error Strategies
   - Panic and Recover
   - Package Error Handling
   - Regular Expressions
9. Appendix - Go Tools

### Advanced
1. Section 1 - Systems Programming
   - Command Line Flags
   - File Descriptors
   - Working with Files
   - Directory Trees
   - Capturing Output
   - Synchronous Commands
   - Asynchronous Commands
   - Pipes with Commands
   - Signal Channels
   - Signal Handlers
   - Resource Management
2. Section 2 - Reflection and Plugins
   - Why Reflection?
   - Type, Value, Kind
   - Structure Fields
   - Structure Tags
   - Modifying Reflection Objects
   - Interface Discovery
   - Reflection Methods
   - What are Plugins?
   - Shared Object Libraries
   - Building Plugins
   - Loading Plugins
   - Plugin Symbols
3. Section 3 - Concurrency
   - Using the sync Package
   - WaitGroup
   - Mutex Locks
   - Concurrent Slices and Maps
   - RWMutex Locks
   - Cond with Signal
   - Cond with Broadcast
   - Using once and atomic
   - Memory Pools
   - Race Conditions
   - Monitors
   - Generators and Coroutines
   - Futures and Pipelines
   - Semaphores
4. Section 4 - Modules
   - Creating Modules
   - Module Files
   - Dependencies
   - Module Testing
   - Semantic Versions
   - Modules with GitHub
   - Versioning and Publishing
   - Dependency Updates
   - Patching Updates
   - Minor Updates
   - Major Updates
   - Multiple Versions
   - Vendor Files
5. Section 5 - Persistence and Databases
   - Package gob
   - Serialization
   - Encoding and Decoding
   - Package sql
   - Database Drivers and Handles
   - Prepared Statements
   - Row Objects
   - Query Results
   - CRUD Operations
   - Transactions and Rollbacks
   - Database Objects
   - Module Files
6. Section 6 - Networking
   - IP Address Protocols
   - Network Interfaces
   - Host Names and Addresses
   - DNS Lookups
   - TCP Clients and Servers
   - Sending and Receiving
   - UDP Clients and Servers
   - UDS Clients and Servers
   - Concurrent Servers
   - Network Objects
   - TCP Server Module
   - TCP Client Module
   - Module Files
7. Section 7 - Web Services
   - JSON Parsing
   - Encode and Decode
   - Marshal and Unmarshal
   - Generic Unmarshals
   - HTTP Web Servers
   - HTTP Handlers
   - Serving from Static Files
   - Serving from Static Directories
   - HTTPS Web Servers
   - RESTful Web Services
   - REST API Endpoints
   - REST Verbs and Status Codes
   - BookCatalog REST Service
   - WebSockets
8. Section 8 - Extending Go
   - Calling C Functions From Go
   - Working with cgo
   - Go and C Conversions
   - Go Module C Libraries
   - Calling Go Functions From C
   - Using export and extern
   - Go Callbacks from C
   - C Shared Libraries
   - Error Handling Strategies
   - Using panic and recover
   - Panics in Modules
9. Appendix - Go Deployments
   - Working with Docker
   - Dockerfiles
   - Go in Containers
   - Multi-Stage Builds
   - Working with Kubernetes
   - Pod Manifests
   - Persistent Volumes
   - Storage Classes
   - Go Program Deployments