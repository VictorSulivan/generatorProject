#include <gtest/gtest.h>
#include "core/calculator.hpp"

using namespace core;

class CalculatorTest : public ::testing::Test {
protected:
    Calculator calc;
};

TEST_F(CalculatorTest, BasicOperations) {
    EXPECT_DOUBLE_EQ(calc.add(2, 3), 5);
    EXPECT_DOUBLE_EQ(calc.subtract(5, 3), 2);
    EXPECT_DOUBLE_EQ(calc.multiply(2, 3), 6);
    EXPECT_DOUBLE_EQ(calc.divide(6, 2), 3);
}

TEST_F(CalculatorTest, AdvancedOperations) {
    EXPECT_DOUBLE_EQ(calc.power(2, 3), 8);
    EXPECT_DOUBLE_EQ(calc.square_root(16), 4);
    EXPECT_DOUBLE_EQ(calc.factorial(5), 120);
}

TEST_F(CalculatorTest, ErrorCases) {
    EXPECT_THROW(calc.divide(1, 0), std::invalid_argument);
    EXPECT_THROW(calc.square_root(-1), std::invalid_argument);
    EXPECT_THROW(calc.factorial(-1), std::invalid_argument);
    EXPECT_THROW(calc.factorial(171), std::overflow_error);
}

TEST_F(CalculatorTest, History) {
    calc.add(1, 2);
    calc.multiply(3, 4);
    
    const auto& history = calc.getHistory();
    EXPECT_EQ(history.size(), 2);
    
    EXPECT_EQ(history[0].type, "add");
    EXPECT_DOUBLE_EQ(history[0].result, 3);
    EXPECT_EQ(history[0].operands.size(), 2);
    
    EXPECT_EQ(history[1].type, "multiply");
    EXPECT_DOUBLE_EQ(history[1].result, 12);
    EXPECT_EQ(history[1].operands.size(), 2);
} 