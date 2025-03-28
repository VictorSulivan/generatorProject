#pragma once

#include <string>
#include <vector>

namespace core {

class Calculator {
public:
    Calculator() = default;
    ~Calculator() = default;

    // Opérations de base
    double add(double a, double b) const;
    double subtract(double a, double b) const;
    double multiply(double a, double b) const;
    double divide(double a, double b) const;

    // Opérations avancées
    double power(double base, double exponent) const;
    double square_root(double number) const;
    double factorial(int number) const;

    // Historique des opérations
    struct Operation {
        std::string type;
        double result;
        std::vector<double> operands;
    };

    void addToHistory(const Operation& operation);
    const std::vector<Operation>& getHistory() const;

private:
    std::vector<Operation> history_;
};

} // namespace core 