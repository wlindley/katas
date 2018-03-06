cmake_minimum_required(VERSION 3.10)
set (CMAKE_CXX_STANDARD 14)

project(PrimeFactors)

include_directories(include)
include_directories(catch)

file(GLOB SOURCES src/*.cpp)
file(GLOB TEST_SOURCES test/*.cpp)

add_executable(PrimeFactors main.cpp ${SOURCES})

add_executable(tests ${TEST_SOURCES} ${SOURCES})
add_test(NAME Tests COMMAND tests)
