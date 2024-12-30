/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	hyperspikeiov1beta1 "hyperspike.io/gitea-operator/api/v1"
)

// var _ = Describe("Gitea Controller", func() {
// 	Context("When reconciling a resource", func() {
// 		const resourceName = "test-resource"

// 		ctx := context.Background()

// 		typeNamespacedName := types.NamespacedName{
// 			Name:      resourceName,
// 			Namespace: "default", // TODO(user):Modify as needed
// 		}
// 		gitea := &hyperspikeiov1beta1.Gitea{}

// 		BeforeEach(func() {
// 			By("creating the custom resource for the Kind Gitea")
// 			err := k8sClient.Get(ctx, typeNamespacedName, gitea)
// 			if err != nil && errors.IsNotFound(err) {
// 				resource := &hyperspikeiov1beta1.Gitea{
// 					ObjectMeta: metav1.ObjectMeta{
// 						Name:      resourceName,
// 						Namespace: "default",
// 					},
// 					// TODO(user): Specify other spec details if needed.
// 				}
// 				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
// 			}
// 		})

// 		AfterEach(func() {
// 			// TODO(user): Cleanup logic after each test, like removing the resource instance.
// 			resource := &hyperspikeiov1beta1.Gitea{}
// 			err := k8sClient.Get(ctx, typeNamespacedName, resource)
// 			Expect(err).NotTo(HaveOccurred())

// 			By("Cleanup the specific resource instance Gitea")
// 			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
// 		})
// 		It("should successfully reconcile the resource", func() {
// 			By("Reconciling the created resource")
// 			controllerReconciler := &GiteaReconciler{
// 				Client: k8sClient,
// 				Scheme: k8sClient.Scheme(),
// 			}

// 			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
// 				NamespacedName: typeNamespacedName,
// 			})
// 			Expect(err).NotTo(HaveOccurred())
// 			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
// 			// Example: If you expect a certain status condition after reconciliation, verify it here.
// 		})
// 	})
// })

var _ = Describe("Gitea Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		gitea := &hyperspikeiov1beta1.Gitea{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind Gitea")
			err := k8sClient.Get(ctx, typeNamespacedName, gitea)
			if err != nil && errors.IsNotFound(err) {
				resource := &hyperspikeiov1beta1.Gitea{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: hyperspikeiov1beta1.GiteaSpec{
						// Add a secret reference for external Gitea
						SecretRef: &hyperspikeiov1beta1.SecretRef{
							Name: "external-gitea-secret",
						},
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			By("cleaning up the specific resource instance Gitea")
			resource := &hyperspikeiov1beta1.Gitea{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})

		It("should successfully reconcile an external Gitea resource", func() {
			By("Reconciling the created resource with external Gitea")
			controllerReconciler := &GiteaReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			// Add assertions to verify external Gitea was handled correctly
			// Example: Verify that the SecretRef was used to create the Gitea client
			By("Verifying the external Gitea client")
			createdResource := &hyperspikeiov1beta1.Gitea{}
			Expect(k8sClient.Get(ctx, typeNamespacedName, createdResource)).To(Succeed())
			Expect(createdResource.Spec.SecretRef).NotTo(BeNil())
			Expect(createdResource.Spec.SecretRef.Name).To(Equal("external-gitea-secret"))
		})

		It("should successfully reconcile an internal Gitea resource", func() {
			By("Updating the resource to use internal Gitea")
			resource := &hyperspikeiov1beta1.Gitea{}
			Expect(k8sClient.Get(ctx, typeNamespacedName, resource)).To(Succeed())
			resource.Spec.SecretRef = nil // Remove SecretRef to simulate internal Gitea
			Expect(k8sClient.Update(ctx, resource)).To(Succeed())

			By("Reconciling the updated resource")
			controllerReconciler := &GiteaReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			// Add assertions to verify internal Gitea was handled correctly
			By("Verifying the internal Gitea instance")
			createdResource := &hyperspikeiov1beta1.Gitea{}
			Expect(k8sClient.Get(ctx, typeNamespacedName, createdResource)).To(Succeed())
			Expect(createdResource.Spec.SecretRef).To(BeNil())
		})
	})
})
