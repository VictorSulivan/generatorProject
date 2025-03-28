use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn test_hello_command() {
    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("hello");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains("Hello, World!"));

    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("hello").arg("--name").arg("Rust");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains("Hello, Rust!"));
}

#[test]
fn test_calc_command() {
    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("calc")
        .arg("--a").arg("5")
        .arg("--b").arg("3")
        .arg("--operation").arg("add");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains("Résultat: 8"));

    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("calc")
        .arg("--a").arg("10")
        .arg("--b").arg("2")
        .arg("--operation").arg("div");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains("Résultat: 5"));
}

#[test]
fn test_file_command() {
    let test_file = "test.txt";
    let test_content = "Hello, World!";

    // Test write
    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("file")
        .arg("--path").arg(test_file)
        .arg("--action").arg("write")
        .arg("--content").arg(test_content);
    cmd.assert().success();

    // Test read
    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("file")
        .arg("--path").arg(test_file)
        .arg("--action").arg("read");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains(test_content));

    // Test delete
    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("file")
        .arg("--path").arg(test_file)
        .arg("--action").arg("delete");
    cmd.assert().success();

    // Verify file is deleted
    assert!(!std::path::Path::new(test_file).exists());
} 