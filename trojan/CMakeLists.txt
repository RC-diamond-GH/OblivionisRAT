cmake_minimum_required(VERSION 3.10)
project(MyProject)

set(CMAKE_CXX_STANDARD 11)

# 添加可执行文件
add_executable(out 
    main.cpp
    Network.hpp
    Network.cpp
    OblivionisAES.hpp
    OblivionisAES.cpp
    base64.hpp
    base64.cpp
    debug.hpp
    debug.cpp
    KeyExchange.hpp
    KeyExchange.cpp
    command.hpp
    command.cpp
)

# 添加链接库
target_link_libraries(out PRIVATE ws2_32 stdc++)