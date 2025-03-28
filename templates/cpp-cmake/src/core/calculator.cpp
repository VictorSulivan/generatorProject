#include "core/calculator.hpp"
#include <cmath>
#include <stdexcept>

namespace core {

double Calculator::add(double a, double b) const {
    double result = a + b;
    addToHistory({"add", result, {a, b}});
    return result;
}

double Calculator::subtract(double a, double b) const {
    double result = a - b;
    addToHistory({"subtract", result, {a, b}});
    return result;
}

double Calculator::multiply(double a, double b) const {
    double result = a * b;
    addToHistory({"multiply", result, {a, b}});
    return result;
}

double Calculator::divide(double a, double b) const {
    if (b == 0) {
        throw std::invalid_argument("Division by zero");
    }
    double result = a / b;
    addToHistory({"divide", result, {a, b}});
    return result;
}

double Calculator::power(double base, double exponent) const {
    double result = std::pow(base, exponent);
    addToHistory({"power", result, {base, exponent}});
    return result;
}

double Calculator::square_root(double number) const {
    if (number < 0) {
        throw std::invalid_argument("Cannot calculate square root of negative number");
    }
    double result = std::sqrt(number);
    addToHistory({"square_root", result, {number}});
    return result;
}

double Calculator::factorial(int number) const {
    if (number < 0) {
        throw std::invalid_argument("Factorial is not defined for negative numbers");
    }
    if (number > 170) {
        throw std::overflow_error("Result would overflow");
    }

    double result = 1;
    for (int i = 2; i <= number; ++i) {
        result *= i;
    }
    addToHistory({"factorial", result, {static_cast<double>(number)}});
    return result;
}

void Calculator::addToHistory(const Operation& operation) {
    history_.push_back(operation);
}

const std::vector<Calculator::Operation>& Calculator::getHistory() const {
    return history_;
}

} // namespace core 